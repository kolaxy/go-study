package main

import "fmt"

func main() {
	fmt.Println(def(1,2,3,4,5))
}

func def(a ... int) int{
	var result int
	for _, i := range a {
		result += i
	}
	return result
}