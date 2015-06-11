### Module 7: Methods & Interfaces

Methods and interfaces in go are quite different from other languages. This is largely because go doesn't have the concept of a class. But as it turns out, go provides a considerable amount of flexibility.

Let's begin with methods. This is pretty interesting because generally methods are function on classes, but in go we don't have the concept of a class. In this case, a method is a function that operates on a named type. This means that we can define a struct, and then define methods on it. Methods can also operate on pointers to named type. This allows us have that method operates on a particular instance of the type.

It be noted that if we want to declare a method on a type that is provided by the language, we need to first name it something else with the `type` keyword.
```go
type Salutations []Salutation

func(salutations Salutations)Greet(do Printer, isFormal bool, times int){
  //...
}
```
