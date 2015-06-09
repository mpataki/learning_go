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
}
