package main

import "fmt"

type Salutation struct {
  name     string
  greeting string
}

func Greet(salutation Salutation) {

}

func main() {
  var s Salutation{"Bob", "Hello"}
  Greet(s)
}
