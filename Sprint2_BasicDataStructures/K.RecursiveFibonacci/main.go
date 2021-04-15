package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		if y > 99999999 { // 9 digits or more
			y %= 1000000000
		}
	}

	pow10 := 1
	for k > 0 {
		pow10 *= 10
		k--
	}

	fmt.Println(x % pow10)
}

// func main() {
// 	var n int
// 	fmt.Scanf("%d", &n)
// 	fmt.Println(fib(n))
// }

// func fib(n int) int {
// 	if n == 0 || n == 1 {
// 		return 1
// 	}

// 	// return fib(n-1) + fib(n-2)
// 	res := fib(n-1) + fib(n-2)
// 	if res > 99999999 { // 9 digits
// 		res %= 1000000000
// 	}

// 	// fmt.Println(res)
// 	return res
// }
