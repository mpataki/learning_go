package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer) {
	message1, message2 := CreateMessage(salutation.name, salutation.greeting, "yo")
	do(message1)
	do(message2)
}

func CreateMessage(name string, greeting ...string) (message1 string, message2 string) {
	greetingLength := len(greeting)
	message1 = greeting[greetingLength-1] + " " + name
	message2 = "HEY! " + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s, Print)
	Greet(s, PrintLine)
}
