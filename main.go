package main

import (
	"fmt"
	"strings"
	"unicode"
	"bufio"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = text[:len(text)-1]
	runedText := []rune(text)
	if unicode.IsUpper(runedText[0]) && strings.HasSuffix(text, ".") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}