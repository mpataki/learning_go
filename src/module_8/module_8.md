### Module 8: Concurrency

Go was designed with currency in mind, and has concurrent constructs built into the language. This makes it pretty unique as in most languages, currency is achieved with the help of an extra library, or through very deliberate attention being given to the problem.

It's often said about go: "Do not communicate by sharing memory. Instead, share memory by communicating".

Concurrency is tackled in go by making one thread the owner of data. If multiple threads need to use this data, it is communicated to other threads via the concept of a channel.

Goroutines in go are basically light weight threads, meaning that they are handled by the Go runtime. By using the keyword `go` in front of a method, we tell the runtime to execute this method concurrently, allowing the existing thread to continue without blocking. Here's an example:
```go
import(
  "time"
  "fmt"
)

func main(){
  go fmt.Fprintf("A")
  fmt.Fprintf("B")
  time.Sleep(100 * time.Millisecond)
}
```

The output from this bit of code is `B` followed by `A`. This is because the printing of `A` is being done by a separate thread that only gets to execute once the first thread has given up the runtime. You'll also notice that we are using the `time` library to synchronize. This is obviously not the ideal way to do this, so let's take a more idiomatic approach to synchronization next.

Channels are created much like slices or maps:
```go
done := make(chan bool)
```
This makes a channel that can pass boolean values, and stores it in done. We use the `<-` operator to read from and write to a channel, like so:
```go
done <- true   // write to the done channel
flag <- done   // read from the done channel
```
When reading from a channel, the operation is blocking if there is nothing on the channel. Therefore, we can use the channel as a method of synchronization.

```go
  done := make(chan bool)

  go func() {
    PrintLine("second")
    done <- true
  }()

  PrintLine("first")
  <- done
```
This declares and defines the `done` channel, defines an anonymous function in-line, and runs it on a goroutine (creating a closure from which it can access the channel). Kicking off the goroutine is a non-blocking action, while reading from the channel is, allowing the main thread to print `"first"` before blocking and waiting for the second thread to finish printing `"second"`, and writing to the channel. The second thread then returns, and the first can wake up as there is now data on the channel for it to read.

It should be noted that our definition of the channel above makes it a non-buffered channel, meaning that it can only hold one item. Consider what would happen if we had written the anonymous function like this instead:
```go
go func() {
  PrintLine("second")
  done <- true
  done <- true
  PrintLine("Never")
}()
```
Here the first `done <- true` would push `true` onto the channel, but the second one would cause the thread to block, waiting for there to be room on the channel to push the additional `true`. In this case, we would never see `"Never"` get printed as the channel isn't read from twice, making the thread block forever.

To make this a buffered channel, we would declare the size of the buffer as follows:
```go
done := make(chan boolean, 2)
```
