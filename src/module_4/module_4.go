package main

import "../greeting"

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
