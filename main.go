package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numString string
	fmt.Scan(&numString)
	for _, v := range numString {
		val, _ := strconv.Atoi(string(v))
		res := int(val) * int(val)
		fmt.Printf("%d", res)
	}
}
