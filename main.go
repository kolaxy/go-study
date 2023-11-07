package main
import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{3, 2, 1}

	fmt.Println(a == b) // true
	fmt.Println(a == c) // false
}