### Module 7: Methods & Interfaces

Methods and interfaces in go are quite different from other languages. This is largely because go doesn't have the concept of a class. But as it turns out, go provides a considerable amount of flexibility.

Let's begin with methods. This is pretty interesting because generally methods are function on classes, but in go we don't have the concept of a class. In this case, a method is a function that operates on a named type. This means that we can define a struct, and then define methods on it. Methods can also operate on pointers to named type. This allows us have that method operates on a particular instance of the type.

It be noted that if we want to declare a method on a type that is provided by the language, we need to first name it something else with the `type` keyword.
```go
type Salutations []Salutation

func(salutations Salutations) Greet(do Printer, isFormal bool, times int){
  //...
}
```

Now let's see an example where we operate on a pointer to a type:
```go
func (salutation *Salutation) Rename(newName string) {
  salutation.Name = newName
}
```
Notice the `*` in front of the type indicating a pointer. This means that the change to the `Name` property on whatever `Salutation` struct that will call this on will be modified, and persistently changed afterwards.

Interfaces are dead simple in go. To begin with, you define what makes up the interface, like so:
```go
type Renamable interface {
  Rename(newName string)
}
```
Implementing the interface just involves defining the methods listed in the interface.
```go
type Salutation struct {
  Name string
}

func (salutation *Salutation) Rename(newName string) {
  salutation.Name = newName
}
```
This is all that's required for a `struct` to implement an interface. Which means we can now do things like this:
```go
func RenameToFrog(r Renamable) {
  r.Rename("Frog")
}

func main() {
  s := Salutation { "MyName" }
  RenameToFrog(&s)
  // now s.Name == "Frog"
}
```

We've used the empty interface in our examples already, like this:
```go
func TypeSwitchTest(x interface{}) {
  // ...
}
```
We know that in order for a type to implement an interface, it must implement the methods listed in the interface. However, because there are no methods to implement in the empty interface, any type implements that interface. We often see this used as we do above, which allows us to pass a variable of any type into the function.

The `Writer` interface is built into go, and looks like this:
```go
type Writer interface {
  Write(p byte[]) (n int, err error)
}
```

We can have our `Salutation` type implement this interface as follows:
```go
func (salutation *Salutation) Write(p []byte) (n int, err error) {
  s := string(p)
  salutation.Rename(s)
  n = len(s)
  err = nil
  return
}
```

Because our `Salutaion` class implements the `Writer` interface, we can pass an instance of it to methods that take in a writer class. For example `fmt.Fprintf`.
```go
import "fmt"

func main() {
  fmt.Fprintf(&salutation, "The count is %d.", 10)
}
```
