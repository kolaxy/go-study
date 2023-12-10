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
	counter := 0
	for i := 0; i < count; i++ {
		if array[i] > 0 {
			counter += 1
		}
	}
	fmt.Println(counter)
}

