package main
import "fmt"

func main() {
	var put float64
	fmt.Scan(&put)
	if put > 10000 {
		fmt.Printf("%e", put)
	} else if put > 0 {
		fmt.Printf("%.4f", put * put)
	} else {
		fmt.Printf("число %.2f не подходит",put)
	}
}
