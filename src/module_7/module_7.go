package main

import "fmt"

type Salutation struct {
  Name     string
  Greeting string
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

func main() {
  salutations := Salutations{
    {"Bob", "Hello"},
    {"Joe", "Hi"},
    {"Merry", "What is up?"},
  }

  salutations.Greet(CreatePrintFunction("?"), true, 6)
}
