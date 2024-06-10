package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)

	if isValidPassword(a) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func isValidPassword(a string) bool {
	for _, symb := range a {
		if !isLatinLetterOrDigit(symb) {
			return false
		}
	}
	return len(a) >= 5
}

func isLatinLetterOrDigit(s rune) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9')
}
