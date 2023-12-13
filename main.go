package main

import (
	"fmt"
	"time"
)

func work(n int) int {
	if n > 3 {
	   time.Sleep(time.Millisecond * 500)
	   return n + 1
	} else {
	   time.Sleep(time.Millisecond * 500)
	   return n - 1
	}
 }
func main() {
	cached := make(map[int]int)
	for i := 0; i < 10; i++ {
		var temp int
		fmt.Scan(&temp)
		val, ok := cached[temp]
		if ok == true {
			fmt.Printf("%v ", val)
		} else {
			loaded := work(temp)
			cached[temp] = loaded
			fmt.Printf("%v ", loaded)
		}
	}
}
