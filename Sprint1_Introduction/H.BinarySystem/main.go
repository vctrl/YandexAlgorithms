package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scanf("%s\n%s", &a, &b)

	var max, min *string
	if len(a) >= len(b) {
		max = &a
		min = &b
	} else {
		max = &b
		min = &a
	}

	result := make([]byte, len(*max))
	var rm byte
	for i := len(*max) - 1; i >= 0; i-- {
		ab := (*max)[i] - "0"[0]

		var bb byte
		if len(*max)-i > len(*min) {
			bb = 0
		} else {
			bb = (*min)[i-(len(*max)-len(*min))] - "0"[0]
		}

		sum := ab + bb + rm
		// 1st byte is going to result
		result[i] = (sum & 1)
		// 2nd byte is remainder
		rm = (sum & 10) >> 1
	}

	if rm > 0 {
		fmt.Print(rm)
	}

	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
	}
}
