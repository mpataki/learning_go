### Module 2: Functions

- Functions in Go can have multiple return values
- We can pass go function just like any other type, similar to javascript
- We can declare function inside of other functions
-- This gives us closure behavior

When declaring a function, much like variables, the name of params comes first, then the type:
```go
func foo(some_number int){
  // ...
}
```

With the following syntax, we are specifying that both params are of type string:
```go
func CreateMessage(name, greeting string) {
  //...
}
```

Notice how above we haven't specified a return type. This is because we aren't returning anything. In other languages we might need `void` for this, however not with go.

We specify the return type of a function as follows:
```go
func CreateMessage(name, greeting string) string {
  return greeting + " " + name
}
```

Returning multiple values from a function:
```go
func CreateMessage(name, greeting string) (string, string) {
  return greeting + " " + name, "HEY! " + name
}

func Greet(salutation Salutation) {
  message1, message2 := CreateMessage("Bob", "Hello")
  fmt.Println(message1)
  fmt.Println(message2)
}
```

Let's say that we didn't want to use one of the two values above:
```go
func Greet(salutation Salutation) {
  message1, message2 := CreateMessage("Bob", "Hello")
  fmt.Println(message1)
} // BAD
```
This thrown an error (that's right, an ERROR) saying that `message2` was declared but not used. We can get around this as follows:
```go
func Greet(salutation Salutation) {
  message1, _ := CreateMessage("Bob", "Hello")
  fmt.Println(message1)
}
```
