package main
import "fmt"

func main() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(numbers[0]) // 1
	fmt.Println(numbers[4]) // 2

	numbers[0] = 87

	fmt.Println(numbers[0]) // 87
}