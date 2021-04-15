package main

import "fmt"

func main() {
	var a, x, b, c int
	fmt.Scanf("%d %d %d %d", &a, &x, &b, &c)
	fmt.Print(a*x*x + b*x + c)
}
