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

Also, you don't need an expression to switch on.

In go, cases can be expressions, not just values like in many other languages.

We can switch on types in go. There is special syntax for this.
