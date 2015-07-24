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

The output from this bit of code is `B` followed by `A`. This is because the printing of `A` is being done by a separate thread that only gets to execute once the first thread has given up the runtime. You'll also notice that we are using the `time` library to synchronize. This is obviously not the ideal way to do this, but we'll get into more idiomatic approaches to synchronization next.
