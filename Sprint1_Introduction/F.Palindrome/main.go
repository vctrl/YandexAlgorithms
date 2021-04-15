package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := strings.ToLower(scanner.Text())

	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		for !unicode.IsLetter(rune(text[i])) && i < j {
			i++
		}
		for !unicode.IsLetter(rune(text[j])) && i < j {
			j--
		}
		if text[i] != text[j] {
			fmt.Print("False")
			return
		}
	}

	fmt.Print("True")
}
