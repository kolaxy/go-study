package main

import "fmt"

func main() {
	minimumFromFour()
}

func minimumFromFour() int {
	array := []int{}
	for i := 0; i < 4; i++ {
		var temp int
		fmt.Scan(&temp)
		array = append(array, temp)
	}
	min := array[0]
	for i := 0; i < 4; i++ {
		if array[i] < min {
			min = array[i]
		}
	}
	return min
}
