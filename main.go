package main

import (
	"fmt"
)

func main() {
	var put string
	fmt.Scan(&put)
	for index, v := range put {
		if index%2 != 0 {
			fmt.Printf("%v", string(v))
		}
	}
}
