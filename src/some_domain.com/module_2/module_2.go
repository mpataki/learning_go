package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	message1, _ := CreateMessage(salutation.name, salutation.greeting, "yo")
	fmt.Println(message1)
}

func CreateMessage(name string, greeting ...string) (message1 string, message2 string) {
	greetingLength := len(greeting)
	message1 = greeting[greetingLength-1] + " " + name
	message2 = "HEY! " + name
	return
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
