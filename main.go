package main

import "fmt"

// !() - nogic NOT
// && () - logic  AND
// || () - logic OR

func main() {
	a := 6
	b := 3
	if a < b {
		fmt.Println("A is less than B")
	} else if a > b {
		fmt.Println("A is greater than B")
	}
}