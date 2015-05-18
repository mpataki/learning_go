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
  do(message1)
}
```
This allows us to execute a statement when gets evaluated before the conditional but that is only scoped within the `if` statement block.
