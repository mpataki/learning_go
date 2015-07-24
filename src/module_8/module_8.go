package main

import (
  "fmt"
)

type Salutation struct {
  Name     string
  Greeting string
}

type Renamable interface {
  Rename(newName string)
}

func RenameToFrog(r Renamable) {
  r.Rename("Frog")
}

type Salutations []Salutation

type Printer func(string)

func(salutations Salutations) Greet(do Printer, isFormal bool, times int){
  for _, s := range salutations {
    message1, message2 := CreateMessage(s.Name, s.Greeting)

    if prefix := GetPrefix(s.Name); isFormal {
      do(prefix + message1)
    } else {
      do(message2)
    }
  }
}

func (salutation *Salutation) Rename(newName string) {
  salutation.Name = newName
}

func GetPrefix(name string) (prefix string) {
  switch name {
  case "Bob":
    prefix = "Mr. "
  case "Joe", "Amy":
    prefix = "Dr. "
  case "Mary":
    prefix = "Mrs. "
  default:
    prefix = "Dude "
  }

  return
}

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

func CreateMessage(name string, greeting string) (message1 string, message2 string) {
  message1 = greeting + " " + name
  message2 = "HEY! " + name
  return
}

func Print(s string) {
  fmt.Print(s)
}

func PrintLine(s string) {
  fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
  return func(s string) {
    fmt.Println(s + custom)
  }
}

func (salutation *Salutation) Write(p []byte) (n int, err error) {
  s := string(p)
  salutation.Rename(s)
  n = len(s)
  err = nil
  return
}

func main() {
  salutations := Salutations{
    {"Bob", "Hello"},
    {"Joe", "Hi"},
    {"Merry", "What is up?"},
  }

  salutations[0].Rename("John")
  // name is now "John"

  RenameToFrog(&salutations[0])
  // Name is now "Frog"

  fmt.Fprintf(&salutations[1], "The count is %d.", 10)

  done := make(chan bool)

  go func() {
    PrintLine("second")
    done <- true
  }()

  PrintLine("first")
  <- done
}
