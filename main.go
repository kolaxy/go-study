package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		array[i] = temp 
	}
	min_array := array[0]
	for _, v := range array {
		if v < min_array {
			min_array = v
		}
	}

	
	count := 0
	for i := 0; i < n; i++ {
		if array[i] == min_array {
			count++
		}
	}
	fmt.Println(count)
}