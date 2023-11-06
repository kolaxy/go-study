package main
import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	a1 := a/100
	a2 := a/10%10
	a3 := a%10
	switch {
		case a1 != a2 && a2 != a3 && a1 != a3:
			fmt.Println("YES")
		case 1 == 1:
			fmt.Println("NO")
		}
}