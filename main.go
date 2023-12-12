package main

import (
	"fmt"
	"math"
)

func calculate(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(calculate(a, b))
}
