package main
import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	output := 0
	for i := 0; i < count; i++ {
		var temp int
		fmt.Scan(&temp)
		if (temp % 8 == 0) && (10 <= temp && temp <= 99){
			output += temp
		}
	}
	fmt.Println(output)
}