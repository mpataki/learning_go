package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	fmt.Println(CreateMessage("Bob", "Hello"))
}

func CreateMessage(name, greeting string) string {
	return greeting + " " + name
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
