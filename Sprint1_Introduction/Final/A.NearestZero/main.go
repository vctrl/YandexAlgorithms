// id посылки 43588919

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	NearestZeros(os.Stdin)
}

func NearestZeros(input *os.File) {
	// fillTestData()
	var n int
	fmt.Scan(&n)

	s := bufio.NewScanner(input)
	b := make([]byte, 0, 10*n)
	s.Buffer(b, 10*n+n)
	s.Scan()

	numbers := make([]int, n)
	numbersStr := strings.Split(s.Text(), " ")

	for i := 0; i < len(numbersStr); i++ {
		nn, _ := strconv.Atoi(numbersStr[i])
		numbers[i] = nn
	}

	prevZeroIndex := -1
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 0 {
			// участок с начала до первого нуля
			if prevZeroIndex == -1 {
				for j := 0; j < i; j++ {
					numbers[j] = i - j
				}
			} else if i-prevZeroIndex > 1 { // все участки между нулями
				d := i - prevZeroIndex
				mid := prevZeroIndex + 1 + d/2
				dist := 0
				// до середины расстояние возрастает
				for j := prevZeroIndex + 1; j < mid; j++ {
					dist++
					numbers[j] = dist
				}

				if d%2 != 1 {
					dist--
				}

				for j := mid; j < i; j++ {
					numbers[j] = dist
					dist--
				}
			}

			prevZeroIndex = i
			continue
		}

		// участок с последнего нуля до конца
		if i == n-1 {
			dist := 1
			for j := prevZeroIndex + 1; j < len(numbers); j++ {
				numbers[j] = dist
				dist++
			}
		}
	}

	var result bytes.Buffer
	result.Grow(n*10 + n)
	for i := 0; i < n; i++ {
		result.WriteString(strconv.Itoa(numbers[i]))
		result.WriteString(" ")
	}

	fmt.Print(result.String())
}
