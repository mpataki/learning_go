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
  message1, message2 := CreateMessage(s.name, s.greeting)
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

We can clean this up a bit by naming our return values:
```go
func CreateMessage(name, greeting string) (message1 string, message2 string) {
  message1 = greeting + " " + name
  message2 = "HEY! " + name
  return
}
```

Variadic functions - We use these to allow multiple parameters to be passed in.
```go
func CreateMessage(name, greeting ...string) (message1 string, message2 string) {
  greetingLength = len(greeting)
  message1 = greeting[greetingLength-1] + " " + name
  message2 = "HEY! " + name
  return
}
```
Notice that we use the built in `len` method to find the length of the given parameter.

Functions in go are their own types. We can pass them as parameters into other function so long as we specify the function signature:
```go
func Greet(salutation Salutation, do func(string)) {
  message1, message2 := CreateMessage(salutation.name, salutation.greeting, "yo")
  do(message1)
  do(message2)
}
```

We can simplify this a bit by making a new type that encapsulates that function signature. Here is a complete example that does this:
```go
package main

import "fmt"

type Salutation struct {
  name     string
  greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer) {
  message1, message2 := CreateMessage(salutation.name, salutation.greeting, "yo")
  do(message1)
  do(message2)
}

func CreateMessage(name string, greeting ...string) (message1 string, message2 string) {
  greetingLength := len(greeting)
  message1 = greeting[greetingLength-1] + " " + name
  message2 = "HEY! " + name
  return
}

func Print(s string) {
  fmt.Print(s)
}

func PrintLine(s string) {
  fmt.Println(s)
}

func main() {
  var s = Salutation{"Bob", "Hello"}
  Greet(s, Print)
  Greet(s, PrintLine)
}
```
