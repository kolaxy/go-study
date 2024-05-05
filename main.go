package main

import "fmt"

func main() {
	m1 := make(map[int]int)
	m2 := map[int]int{
		12: 2,
		1:  5,
	}
	fmt.Println(m1)
	fmt.Println(m2)
}
