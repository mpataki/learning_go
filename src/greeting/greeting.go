package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message1, message2 := CreateMessage(salutation.Name, salutation.Greeting)

	if prefix := "Mr. "; isFormal {
		do(message1)
	}

	do(message2)
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
