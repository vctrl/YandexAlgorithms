package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 {
		fmt.Print("True")
		return
	}

	result := next(1, n)
	if result {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}

func next(fourPow int, n int) bool {
	if fourPow*4 == n {
		return true
	}
	if fourPow*4 > n {
		return false
	}

	return next(fourPow*4, n)
}
