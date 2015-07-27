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

func (salutations Salutations) ChannelGreet(c chan Salutation) {
  for _, s := range salutations {
    c <- s
  }
  close(c)
}

func main() {
  salutations := Salutations{
    {"Bob", "Hello"},
    {"Joe", "Hi"},
    {"Merry", "What is up?"},
  }

  done := make(chan bool)

  go func() {
    PrintLine("second")
    done <- true
  }()

  PrintLine("first")
  <- done

  c := make(chan Salutation, 10)
  salutations.ChannelGreet(c)

  for s := range c {
    PrintLine(s.Name)
  }

  // now for a select example
  ch1 := make(chan Salutation, 10)
  ch2 := make(chan Salutation, 10)

  go salutations.ChannelGreet(ch1)
  go salutations.ChannelGreet(ch2)

  for {
    select {
      case s, ok := <- ch1:
        if ok {
          fmt.Println(s, ":1")
        } else {
          return
        }
      case s, ok := <- ch2:
        if ok {
          fmt.Println(s, ":2")
        } else {
          return
        }
      default:
        fmt.Println("Waiting...")
    }
  }
}
