package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_INT = 10001
	MIN_INT = -1
)

func main() {
	var n, m int
	fmt.Scanf("%d\n%d", &n, &m)

	if n == 0 && m == 0 {
		return
	}

	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 10000*20)
	s.Scan()
	nstr := strings.Split(strings.TrimSpace(s.Text()), " ")
	s.Scan()
	mstr := strings.Split(strings.TrimSpace(s.Text()), " ")

	ns := make([]int, n)
	ms := make([]int, m)

	for i := 0; i < n; i++ {
		ns[i], _ = strconv.Atoi(nstr[i])
	}
	for i := 0; i < m; i++ {
		ms[i], _ = strconv.Atoi(mstr[i])
	}

	median := findMedian(ns, ms, n, m)

	fmt.Print(median)
}

func findMedian(ns, ms []int, n, m int) float64 {
	var median float64

	if m == 0 {
		if (n+m)%2 == 0 {
			median = float64(ns[len(ns)/2]+ns[len(ns)/2-1]) / 2

			return median
		}

		median = float64(ns[(len(ns)+1)/2-1])
		return median
	}
	if n == 0 {
		if (n+m)%2 == 0 {
			median = float64(ms[len(ms)/2]+ms[len(ms)/2-1]) / 2
			return median
		}

		median = float64(ms[(len(ms)+1)/2-1])
		return median

	}

	big, sm := getBiggerSmaller(ns, ms)

	low := 0
	high := len(sm)
	var x, y int

	var leftX, rightX int
	var leftY, rightY int

	for {
		x = ((low + high) + 1) / 2
		y = (len(big)+len(sm))/2 - x

		// fmt.Println(x, y)

		if x <= 0 {
			leftX = MIN_INT
		} else {
			leftX = sm[x-1]
		}

		if y <= 0 {
			leftY = MIN_INT
		} else {
			leftY = big[y-1]
		}

		if x >= len(sm) {
			rightX = MAX_INT
		} else {
			rightX = sm[x]
		}

		if y >= len(big) {
			rightY = MAX_INT
		} else {
			rightY = big[y]
		}

		// fmt.Println(leftX, rightY, leftY, rightX)
		if leftX <= rightY && leftY <= rightX {
			var maxLeft, minRight int
			if leftX >= leftY {
				maxLeft = leftX
			} else {
				maxLeft = leftY
			}

			if rightX <= rightY {
				minRight = rightX
			} else {
				minRight = rightY
			}

			if (n+m)%2 == 0 {
				median = float64(maxLeft+minRight) / 2
			} else {
				median = float64(minRight)
			}

			break
		}

		if rightY > leftX {
			low = x
		}

		if rightY < leftX {
			high = x - 1
		}
	}

	return median
}

func getBiggerSmaller(ns, ms []int) ([]int, []int) {
	if len(ns) >= len(ms) {
		return ns, ms
	}
	return ms, ns
}
