package main

import "fmt"

func main() {
	fmt.Println(sumInt(1,2,3,4,5))
}

func sumInt(a ... int) (int, int){
	var result int
	var count int
	for _, i := range a {
		result += i
		count++
	}
	return count, result
}