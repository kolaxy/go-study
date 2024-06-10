package main

import "fmt"

func main() {
	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)

	fmt.Println(bigIndex)
	fmt.Printf("%T \n", bigIndex)

	var a int32 = 22
	var b uint64
	b = uint64(a)

	fmt.Println(b)
	fmt.Printf("%T \n", b)
}
