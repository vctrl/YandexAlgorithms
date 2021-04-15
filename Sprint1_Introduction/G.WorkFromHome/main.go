package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Print(0)
	}

	result := make([]int, 0)

	result = next(n, result)

	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
}

func next(n int, res []int) []int {
	if n == 0 {
		return res
	}

	res = append(res, n%2)
	return next(n/2, res)
}
