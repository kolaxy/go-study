package main
import "fmt"

func main() {
	maxelement := 0
	counter := 0
	for i := 0; i <= 1; {
		var put int
		fmt.Scan(&put)
		if put == 0 {
			break
		} else if put == maxelement {
			counter++
		} else if put > maxelement {
			counter=1
			maxelement = put
		}
	}
	fmt.Println(counter)
}