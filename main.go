package main
import "fmt"

func main() {
	var numcount int
	counter := 0
	fmt.Scan(&numcount)
	for i := 0; i < numcount; i++ {
		var temp int
		fmt.Scan(&temp)
		if temp == 0 {
			counter++
		}
	}
	fmt.Println(counter)
}