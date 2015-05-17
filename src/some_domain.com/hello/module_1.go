package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

const (
	PI       = 3.14
	Language = "Go"
)

const (
	A = iota
	B = iota
	C = iota
)

const (
	D = iota
	E
	F
)

func main() {
	message := "hello, world"
	greeting := &message

	fmt.Println(greeting, *greeting)

	*greeting = "something else"
	fmt.Println(greeting, *greeting)

	var s1 Salutation = Salutation{"guy", "hello!"}
	fmt.Println(s1.name)
	fmt.Println(s1.greeting)

	var s2 Salutation = Salutation{greeting: "hello!", name: "guy"}
	fmt.Println(s2.name)
	fmt.Println(s2.greeting)

	var s3 Salutation = Salutation{}
	s3.name = "guy"
	s3.greeting = "greetings!"
	fmt.Println(s3.name)
	fmt.Println(s3.greeting)

	fmt.Println(PI)
	fmt.Println(Language)

	fmt.Println(A, B, C)
	fmt.Println(D, E, F)
}
