package main

import "fmt"

func main() {
	var put string
	fmt.Scan(&put)
	for index, line := range put {
		if index == len(put)-1 {
			fmt.Printf("%v", string(line))
		} else {
			fmt.Printf("%v*", string(line))
		}
	}
}
