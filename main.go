package main
import (
	"fmt"
	"strconv"
)

func calculate_sum(str string) (result int){
	for _, letter := range str {
		num, err := strconv.Atoi(string(letter))
		if err == nil {
			result += num
		}
	}
	return 
}

func main() {
	var ticker_number string
	fmt.Scan(&ticker_number)
	first := string(ticker_number[:3])
	second := string(ticker_number[3:])
	res := (bool(calculate_sum(first) == calculate_sum(second)))
	if res == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}