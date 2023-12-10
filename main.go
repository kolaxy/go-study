package main
import "fmt"

func main() {
	var count int
	var ls []int
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		var temp int
		fmt.Scan(&temp)
		ls = append(ls, temp)
	}
	fmt.Println(ls[3])
}

