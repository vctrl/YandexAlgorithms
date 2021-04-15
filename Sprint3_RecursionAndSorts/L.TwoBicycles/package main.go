package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	s := bufio.NewScanner(os.Stdin)

	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 1000000*(n+1))
	s.Scan()
	a := strings.Split(s.Text(), " ")

	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	if n == 1 && k == 1 {
		fmt.Print("0")
		return
	}

	var b bool
	if k < n*(n-1)/2-k {
		b = true
		// dn = k
	} else {
		b = false
		// dn = n*(n-1)/2 - k
	}
	// dn := 5

	dn := n * (n - 1) / 4

	diffs := make([]int, dn)
	var m int

	if b {
		m = dn - 1
	} else {
		m = 0
	}

	first := true
	fmt.Println("asdf")
	for i := 0; i < len(a); i++ {
		fmt.Printf("s%d: %v\n", i, diffs)
		a1, _ := strconv.Atoi(a[i])

		if b {
			if first || a1 <= diffs[m] {
				fmt.Printf("%d <= %d\n", a1, diffs[m])
				diffs[m] = a1
				m--
			}

			if m < 0 {
				first = false
				sort.Ints(diffs)
				m = dn - 1
			}
		} else {
			if first || a1 >= diffs[m] {
				fmt.Printf("%d <= %d, m = %d\n", a1, diffs[m], m)
				diffs[m] = a1
				m++
			}

			if m == dn {
				first = false
				sort.Ints(diffs)
				m = 0
			}
		}

	}

	sort.Ints(diffs)

	if b {
		fmt.Println(diffs[k-1])
	} else {
		fmt.Println(diffs[dn-k-1])
	}

	// fmt.Println(diffs)
	// просто допустим что k в первой половине
	// dn := n * (n + 1) / 4
	// diffs := make([]int, dn)
	// fmt.Println(len(diffs))

	// m := dn -1
	// first := true
	// for i := 0; i < len(a); i++ {
	// 	for j := i + 1; j < len(a); j++ {
	// 		a1, _ := strconv.Atoi(a[j])
	// 		a2, _ := strconv.Atoi(a[i])

	// 		if first || mod(a1-a2) <= diffs[dn-1] {
	// 			diffs[m] = mod(a1 - a2)
	// 			m--
	// 		}

	// 		if m < 0 {
	// 			first = false
	// 			sort.Ints(diffs)
	// 			m = dn - 1
	// 		}
	// 	}
	// }

	// sort.Ints(diffs)
	// fmt.Print(diffs[k-1])
}

func mod(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}
