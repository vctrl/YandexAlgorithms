package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if a&1 == b&1 && b&1 == c&1 {
		fmt.Print("WIN")
	} else {
		fmt.Print("FAIL")
	}
}
