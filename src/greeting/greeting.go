package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message1, message2 := CreateMessage(salutation.Name, salutation.Greeting)

	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message1)
	} else {
		do(message2)
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
