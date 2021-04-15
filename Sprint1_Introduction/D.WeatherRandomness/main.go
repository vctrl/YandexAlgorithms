package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, 4096)
	scanner.Buffer(buf, 100000*5) // 4 bytes per value + spaces

	scanner.Scan()
	text := scanner.Text()
	numbersStr := strings.Split(text, " ")
	numbers := make([]int16, len(numbersStr))
	for i := 0; i < n; i++ {
		ni, _ := strconv.Atoi(numbersStr[i])
		numbers[i] = int16(ni)
	}

	var result int

	if n == 1 {
		result = 1
	}
	if n > 1 && numbers[0] > numbers[1] {
		result++
	}
	if n > 1 && numbers[n-1] > numbers[n-2] {
		result++
	}
	for i := 1; i < n-1; i++ {
		if numbers[i] > numbers[i-1] && numbers[i] > numbers[i+1] {
			result++
		}
	}

	fmt.Print(result)
}
