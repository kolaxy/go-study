package main

import (
	"fmt"
	// "strings"
	// "unicode"
	"bufio"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	flag := true
	textRune := []rune(text)
	for i := 0; i < len(textRune); i++ {
		if textRune[i] != textRune[len(textRune)-i-1] {
			flag = false
		}
	}
	if flag == true {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
	// text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// text = text[:len(text)-1]
	// runedText := []rune(text)
	// if unicode.IsUpper(runedText[0]) && strings.HasSuffix(text, ".") {
	// 	fmt.Println("Right")
	// } else {
	// 	fmt.Println("Wrong")
	// }
}
