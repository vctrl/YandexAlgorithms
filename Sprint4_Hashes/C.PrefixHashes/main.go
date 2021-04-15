package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, m int
	fmt.Scanf("%d\n%d", &a, &m)
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	sc.Buffer(buf, 10e6)

	sc.Scan()
	str := sc.Text()
	pre_h := prehashes(str, a, m)
	ps := pows(len(str), a, m)

	// fmt.Println(pre_h)
	// fmt.Println(ps)

	var res strings.Builder
	// skip rows quantity
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for sc.Scan() && n > 0 {
		args := strings.Split(sc.Text(), " ")
		l, _ := strconv.Atoi(args[0])
		r, _ := strconv.Atoi(args[1])

		h_ab := pre_h[r-1]
		var h_a int
		if l == 1 {
			h_a = 0
		} else {
			h_a = pre_h[l-2]
		}
		// fmt.Println("h_ab", h_ab, "h_a", h_a, ps[r-l+1])
		// fmt.Println(h_ab - (h_a*ps[r-l+1])%m)

		// h_b := (h_ab - (h_a*ps[r-l+1])%m) % m

		h_b := ((h_ab - (h_a*ps[r-l+1])%m) + m) % m
		// if h_b < 0 {
		// 	q := h_b/m - 1
		// 	h_b = h_b - (q * m)
		// } else {
		// 	h_b = h_b % m
		// }

		// fmt.Println(h_b)
		res.WriteString(strconv.Itoa(h_b))
		if n != 1 {
			res.WriteString("\n")
		}
		n--
	}

	fmt.Println(res.String())
}

func prehashes(str string, q, R int) []int {
	pre_h := make([]int, len(str))

	var h int

	h = int(str[0])
	pre_h[0] = h

	for i := 1; i < len(str); i++ {
		pre_h[i] = (pre_h[i-1]*q + int(str[i])) % R
	}

	return pre_h
}

func pows(n, b, m int) []int {
	res := make([]int, n+1)
	res[0] = 1
	for i := 1; i <= n; i++ {
		r := res[i-1] * b
		res[i] = r % m
	}

	return res
}
