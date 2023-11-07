package main
import "fmt"

func main() {
	var a float64 = 100.123456
	fmt.Printf("This number %f type %T", a, a)
	var a1 byte = 's'
	var a2 int = 1234
	fmt.Printf("%q %b", a1, a2)
	var b1 string = "123"
	var b2 string = "1234"
	fmt.Printf("%q \n%s", b1, b2)
}