package main

import "fmt"

func main() {
	var numString string
	var maxNumber uint8
	fmt.Scan(&numString)
	maxNumber = uint8(numString[0])
	for _, v := range numString {
		if byte(v) > maxNumber {
			maxNumber = byte(v)
		}
	}
	fmt.Println(string(maxNumber))
}
