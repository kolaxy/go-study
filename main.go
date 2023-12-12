package main

import (
	"fmt"
	"strings"
)

func main() {
	var put string
	fmt.Scan(&put)
	for _, v := range put {
		if strings.Count(put, string(v)) == 1 {
			fmt.Printf("%v", string(v))
		}
	}
}
