package main

import (
	"fmt"
)

func main() {
	m := map[int]int{
		12: 2,
		1:  5,
	}
	delete(m, 12)
	if value, inMap := m[1]; inMap {
		fmt.Println(value) // 10
	}
	
	if value, inMap := m[2]; inMap {
		fmt.Println(value) // Условие не выполняется
	}
	fmt.Println(m[12])
	fmt.Println(m[1])
}
