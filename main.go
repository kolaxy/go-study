package main

import (
	// "errors"
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	return a / b, nil
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result, error := divide(a, b)
	if error == nil {
		fmt.Println(result)
	} else {
		fmt.Println(errors.New("ошибка"))
	}
}
