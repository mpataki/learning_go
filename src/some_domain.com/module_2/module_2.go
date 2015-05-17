package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	message1, _ := CreateMessage("Bob", "Hello")
	fmt.Println(message1)
}

func CreateMessage(name, greeting string) (string, string) {
	return greeting + " " + name, "HEY! " + name
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
