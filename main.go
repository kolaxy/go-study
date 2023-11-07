package main
import "fmt"

func main() {
	var x, p, y int
	years := 0
	// x - start point
	// p - percentage point
	// y - finish point
	fmt.Scan(&x, &p, &y)
	for i := 0; i <= 1; {
		x = int(x + (x * p / 100))
		years++
		if x > y {
			fmt.Println(years)
			break
		}
	}
}