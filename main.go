package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		var temp uint8
		fmt.Scan(&temp)
		workArray[i] = temp
	}

	for i := 0; i < 3; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
	}
	
	for _, element := range workArray {
		fmt.Printf("%v ", element)
	}
}