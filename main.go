package main
import "fmt"

func main() {
	temp := 0
	for i := 0; i <= 1; {
		fmt.Scan(&temp)
		if temp < 10 {
			continue
		} else if temp > 100 {
			break
		} else {
			fmt.Println(temp)
		}

	}
}