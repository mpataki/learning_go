package main

import "fmt"

func GetPrefix(name string) (prefix string) {
	var prefixMap map[string]string
	prefixMap = make(map[string]string)

	prefixMap["Bob"] = "Mr."
	prefixMap["Joe"] = "Dr."
	prefixMap["Amy"] = "Dr."
	prefixMap["Mary"] = "Ms."

	return prefixMap[name]
}

func main() {
	name := "Joe"
	fmt.Println(GetPrefix(name), name)
}
