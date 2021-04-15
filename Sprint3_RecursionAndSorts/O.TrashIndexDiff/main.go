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
	str := strings.Split(s.Text(), " ")
	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	if n == 1 && k == 1 {
		fmt.Print("0")
		return
	}

	a := make([]int, n)
	for i := 0; i < len(str); i++ {
		a[i], _ = strconv.Atoi(str[i])
	}

	sort.Ints(a)

	diffs := make([]int, 1000000)

	k1 := k
LOOP:
	for step := 1; step < len(a); step++ {
		for i := 0; i < len(a)-step; i++ {
			d := mod(a[i+step] - a[i])
			diffs[d]++
			k1++
			if k1 == k {
				break LOOP
			}
		}
	}

	i := -1
	for k > 0 {
		i++
		k -= diffs[i]
	}

	fmt.Print(i)
}

func mod(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}
