package main

import "fmt"

func main() {
	var num string
	var digit string
	fmt.Scan(&num, &digit)

	for _, v := range num {
		if string(v) != digit {
			fmt.Printf("%c", v)
		}
	}
}
