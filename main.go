package main

import "fmt"

func main() {
	x1 := 1
	x2 := 2
	test(&x1, &x2)
}

// считайте что fmt уже импортирован и main объявлен
func test(x1 *int, x2 *int) {
	// здесь пишите ваш код
	temp := *x1
	*x1 = *x2
	*x2 = temp
	fmt.Println(*x1, *x2)
}
