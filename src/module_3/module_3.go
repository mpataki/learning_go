package main

import "learning_go/src/greeting"

func main() {
	var s = greeting.Salutation{"Joe", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), true, 1)
	greeting.TypeSwitchTest(1)
}
