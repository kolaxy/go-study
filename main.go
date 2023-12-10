package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	result := 1 + ((num - 1) % 9)
	fmt.Print(result)
}