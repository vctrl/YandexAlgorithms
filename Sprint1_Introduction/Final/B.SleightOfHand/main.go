package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	var k int
	fmt.Scan(&k)
	numbers := make([]int, 10)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		for i := 0; i < utf8.RuneCountInString(line); i++ {
			if line[i] == "."[0] {
				continue
			}

			n, _ := strconv.Atoi(string(line[i]))
			numbers[n]++
		}
	}

	var wins int
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != 0 && k*2 >= numbers[i] {
			wins++
		}
	}

	fmt.Print(wins)
}
