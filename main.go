package main

import "fmt"

func main() {
	var age, name = add(4,5,"Tom", "Simpson")
	fmt.Println(age)
	fmt.Println(name)
}

func add(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}
