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
