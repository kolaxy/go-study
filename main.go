package main

import "fmt"

func main() {
	age := 26
	name := "nv"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// PrintLN
	fmt.Println("hello world!")
	fmt.Println("goodbye world!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings)
	fmt.Printf("my age is %v and my name is %v\n", age, name)
	fmt.Printf("my age is %q and my name is %q\n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %f points\n", 255.55)
	fmt.Printf("you scored %0.1f points\n", 255.55)

	// Sprintg (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println("the saved string is:", str)
}
