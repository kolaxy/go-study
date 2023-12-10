package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var temp int
	flag := false
	for i := a; i <= b; i++ {
		if i % 7 == 0 {
			temp = i
			flag = true
		}
	}
	if flag == true {
		fmt.Println(temp)
	} else {
		fmt.Println("NO")
	}
}
