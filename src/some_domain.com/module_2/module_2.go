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

func CreateMessage(name, greeting string) (message1 string, message2 string) {
	message1 = greeting + " " + name
	message2 = "HEY! " + name
	return
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
