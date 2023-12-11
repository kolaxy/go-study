package main

import "fmt"

func main() {
	a := vote(1, 1 ,1)
	fmt.Println(a)
}

func vote(x int, y int, z int) int {
    if x == y || x == z {
		return x
	} else {
		return y
	}
}





