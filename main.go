package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for idx, num := range a {
		fmt.Printf("Element with %d index : %d\n", idx, num)
	}
}