package main
import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	hours := number / 3600
	minutes := (number - hours * 3600) / 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}