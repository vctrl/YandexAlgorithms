package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)

	buf := make([]byte, 0, n+1)
	scanner.Buffer(buf, n+1)

	scanner.Scan()
	words := strings.Split(scanner.Text(), " ")
	var longest string
	for i := 0; i < len(words); i++ {
		if len(words[i]) > len(longest) {
			longest = words[i]
		}
	}

	fmt.Println(longest)
	fmt.Print(len(longest))
}
