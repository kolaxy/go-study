package main
import (
	"fmt"
	"strconv"
)

func main() {
	var a int64
	fmt.Scan(&a)
	result := strconv.FormatInt(a, 2)
	fmt.Println(result)
}