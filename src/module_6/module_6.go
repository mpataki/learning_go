package main

import "fmt"

func main() {

  // long form initialization
	var s []int
	s = make([]int, 3, 10)
  s[0] = 1
  s[1] = 2
  s[2] = 3

  // short form
  s2 := []int { 1, 2, 3 }

  fmt.Println(s2[:1]) // will print [1]

  s2 = append(s2[:1], s2[2:]...)

  fmt.Println(s2) // will print [1, 3]
}
