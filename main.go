package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		temp := math.Pow(2, float64(i))
		if temp <= float64(num) {
			fmt.Printf("%d ", int(temp))
		} else {
			break
		}
	}
}