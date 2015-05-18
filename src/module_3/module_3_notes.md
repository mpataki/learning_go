### Module 2: Branching

`if` statements work pretty much the same as in other languages. However there are a couple of things to notice. The `{}` braces are mandatory, and there are no parenthesis. Ex:
```go
if isFormal {
  do(message1)
}
```

We can also have an optional statement before the condition.
```go
if prefix := "Mr. "; isFormal {
  do(prefix + message1)
}
```
This allows us to specify a statement that gets evaluated before the conditional but that is only scoped within the `if` statement block. You can even use the statement as a part of the `if` conditional.

`else` is also pretty much the same as other languages.
```go
if prefix := "Mr. "; isFormal {
  do(prefix + message1)
} else {
  do(message2)
}
```

`switch` statements are slightly different in go. The first big difference is that there is no default fall through - so no `break` statements, unless we tell it to fall through explicitly.
```go
switch name {
  case "Bob": prefix = "Mr. "
  case "Joe": prefix = "Dr. "
  case "Mary": prefix = "Mrs. "
  default: prefix = "Dude "
}
```
Or with the `fallthrough`:
```go
switch name {
  case "Bob": prefix = "Mr. "
  case "Joe": prefix = "Dr. "
  case "Mary":
    prefix = "Mrs. "
    fallthrough
  default: prefix = "Dude "
}
```

You can also use lists inside the case statement so that a case can happen under multiple conditions.
```go
switch name {
  case "Bob": prefix = "Mr. "
  case "Joe", "Amy": prefix = "Dr. "
  case "Mary": prefix = "Mrs. "
  default: prefix = "Dude "
}
```

Also, you don't need an expression after the switch, and you can use expressions in the case statements.
```go
switch {
case name == "Bob":
  prefix = "Mr. "
case name == "Joe", name == "Amy", len(name) == 10:
  prefix = "Dr. "
case name == "Mary":
  prefix = "Mrs. "
default:
  prefix = "Dude "
}
```

We can switch on types in go. There is special syntax for this.
```go
func TypeSwitchTest(x interface{}) {
  switch x.(type) {
  case int:
    fmt.Println("int")
  case string:
    fmt.Println("string")
  case Salutation:
    fmt.Println("salutation")
  default:
    fmt.Println("unknown")
  }
}
```
So there is a fair amount of new stuff happening here. Not all of it will be covered now but just to give a quick overview, the type of `x` is an `interface{}` which is akin to a `void*` where we aren't specifying an explicit type. Really, this is the base type, however we won't get into that here. The point of using it is to show that we can call `x.(type)` in order to pull the type out of `x` and then switch on it. Kind of cool.
