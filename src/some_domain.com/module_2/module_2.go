package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	message1, message2 := CreateMessage("Bob", "Hello")
	fmt.Println(message1)
	fmt.Println(message2)
}

func CreateMessage(name, greeting string) (string, string) {
	return greeting + " " + name, "HEY! " + name
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
