package main

import "../greeting"

func main() {
	var s = greeting.Salutation{"Joe", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), true)
}
