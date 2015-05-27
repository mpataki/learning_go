### Module 4: Loops

There is only one looping construct in go - `for`. However, we can create the equivalent of many of common looping constructs with a couple of key words.

```go
for {
  // ...
}
```
This is an infinite loop that you would probably `break` out of explicitly.

```go
for i := 0; i < 5 ; i++ {
  // ...
}
```
Here we have a basic for loop.

```go
for i < 5 {
  // ...
}
```
Here we have something like a `while` loop.

We also having a `continue` statement that works the same as many other languages.
```go
for i := 0; i < 10; i++ {
  if i % 2 == 0 {
    continue
  }

  fmt.Println("got here!")
}
```
Here `"got here!"` would only print 6 times.

We also have a `range` keyword that will help us loop through a collection. `range` will work over arrays, slices, strings, maps, and channels. (In face with channel, this will be a blocking loop as it waits for input on the channel - pretty cool)
```go
for i, s := range slice {
  // ...
}
```
Here `i` is the index and `s` is the value in the slice that we are looping over. If we don't use the index, we need to change it to the following in order for the compiler not to complain:
```go
for _, s := range slice {
  // ...
}
```
