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

