### Module 1: Variables, Types, and Pointers

This stuff is pretty straight forward. In general we declare a variable with:
```go
var variable_name type
```
So the `var` to indicate that we are declaring a new variable. We then follow it by the variable name, then the type.

Go can also infer the type if you are initializing on the same line. For example:
```go
var message = "this is a message"
```

We can also get rid of var, if we use the `:=` operator, which declares and assigns all at one.
```go
message := "this is a message"
```

You can declare and assign multiple variables as once:
```go
a, b, c := 1, 2, 3
```

You can define your own types using:
```go
type NewString string
```

This sets the new type equal to the type `string`. We can use structs to make this more interesting:
```go
type NewType struct{
  some_string string
  some_int int
}
```

Then we can use this struct as follows:
```go
var nt NewType{"blah", 52}
```

This initializes the fields in the struct in the order that they were declared. We can also do:
```go
var nt NewType{some_int: 52, some_string: "blah"}
// or
var nt NewType
nt.some_int = 52
nt.some_string = "blah"
```

Constant types are also available:
```go
const PI = 3.14
```

If we want to declare many `const`s we can do:
```go
const(
  PI = 3.14
  Language = "Go"
)
```

In general we can rely on the types of `const`s to be determined by the language based on how they are assigned and used.

`iota` is another interesting feature of the language. It allows us to simplify the counting process when declaring a wack of `const` variables that we just want to have incrementing values, similar C++'s `enum`
```go
const (
  A = iota
  B = iota
  C = iota
)
```
This will yield the values `A = 1`, `B = 2`, `C = 3`. `const` values will we carried forward from previous assignments if none are specified, letting us do stuff like this:
```go
const (
  A = iota
  B
  B
)
```
This will yield the same result as the previous example.
