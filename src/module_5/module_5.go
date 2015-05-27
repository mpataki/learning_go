package main

import "fmt"

func GetPrefix(name string) (prefix string) {
	prefixMap := map[string]string{
		"Bob":  "Mr.",
		"Joe":  "Dr.",
		"Amy":  "Dr.",
		"Mary": "Ms.",
	}

	return prefixMap[name]
}

func main() {
	name := "Joe"
	fmt.Println(GetPrefix(name), name)
}
