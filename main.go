package main
import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	array := make([]int, count)
	for i := 0; i < count; i++ {
		var temp int
		fmt.Scan(&temp)
		array[i] = temp
	}
	for i := 0; i < count; i++ {
		if i % 2 == 0 {
			fmt.Printf("%d ", array[i])
		}
	}
}

