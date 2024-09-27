package main

import "fmt"

func main() {
	// string
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	// no outside the function!
	nameFour := "yoshi"

	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	//var num0ne int8 = 28
	//var numTwo int8 = -128
	//var numThree uint16 = 256

	//var scoreOne float32 = -25.98
	//var scoreTwo float64 = 342222222235465365436524625654265.8
	//scoreThree := 1.5

}
