package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	
	for idx := range a {
		fmt.Println(a[idx])
	}

	for idx, _ := range a {
		fmt.Println(a[idx])
	}

	for _, elem := range a {
		fmt.Println(elem)
	}
}