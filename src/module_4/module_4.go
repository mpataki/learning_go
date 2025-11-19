package main

import "learning_go/src/greeting"

func main() {
	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Merry", "What is up?"},
	}

	for _, s := range slice {
		greeting.Greet(s, greeting.CreatePrintFunction("!!!"), true, 5)
	}
}
