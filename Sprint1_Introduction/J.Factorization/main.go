package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	lp := make([]int, 0)
	primes := make([]int, 0)

	factorization := make([]int, 0)

	var sqrt int
	for i := 2; i*i <= n; i++ {
		if i >= len(lp) {
			zeros := make([]int, i-len(lp)+1)
			lp = append(lp, zeros...)
		}

		if lp[i] == 0 {
			lp[i] = i

			for n%i == 0 {
				n /= i
				sqrt = int(math.Sqrt(float64(n)))
				factorization = append(factorization, i)
				if n == 1 {
					fmt.Print(strings.Trim(fmt.Sprint(factorization), "[]"))
					return
				}
			}

			primes = append(primes, i)
			for j := 0; j < len(primes); j++ {
				x := i * primes[j]
				if primes[j] > lp[i] || x > n {
					break
				}

				if x >= len(lp) {
					if x > sqrt {
						continue
					} else {
						zeros := make([]int, sqrt-len(lp)+1)
						lp = append(lp, zeros...)
					}
				}

				lp[x] = primes[j]
			}
		}
	}

	factorization = append(factorization, n)
	fmt.Print(strings.Trim(fmt.Sprint(factorization), "[]"))
}
