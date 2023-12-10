package main
import "fmt"

func main()  {
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
	temp := array[0]
	for i:=0; i < len(array); i++ {
		if array[i] > temp {
			temp = array[i]
		}
	}
	fmt.Println(temp)
}