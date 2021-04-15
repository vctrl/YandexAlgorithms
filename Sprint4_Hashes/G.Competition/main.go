package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 10e5*2+2)

	s.Scan()
	s.Scan()
	input := s.Text()
	if input == "" {
		fmt.Print(0)
		return
	}
	str := strings.Split(input, " ")
	// 1st is 1 or 0
	// 2nd is distance
	// 3rd are indexes
	ranges := make(map[int]map[int][]int)
	ranges[0] = make(map[int][]int)
	ranges[1] = make(map[int][]int)

	curr := 0
	for i, b := range str {
		switch b {
		case "0":
			curr--
			if _, ok := ranges[0][curr]; !ok {
				ranges[0][curr] = make([]int, 0, 1)
			}

			ranges[0][curr] = append(ranges[0][curr], i)
		case "1":
			curr++
			if _, ok := ranges[1][curr]; !ok {
				ranges[1][curr] = make([]int, 0, 1)
			}

			ranges[1][curr] = append(ranges[1][curr], i)
		}

	}

	// fmt.Println(ranges)
	max := 0
	currMax := 0
	for d, i1 := range ranges[0] {
		// d + 1
		if i2, ok := ranges[1][d+1]; ok {
			// fmt.Println(i2, i1, i2[len(i2)-1], i1[0])
			currMax = i2[len(i2)-1] - i1[0]
			if currMax < 0 {
				currMax = -currMax
			}
			// fmt.Println("adsf", currMax)
			// m1 := i2[len(i2)-1] - i1[0]
			// m2 := i1[len(i1)-1] - i2[0]
			// if m1 > m2 {
			// 	currMax = m1
			// } else {
			// 	currMax = m2
			// }
		}
		if currMax > max {
			max = currMax
		}
		if i2, ok := ranges[0][d+1]; ok {
			// fmt.Println(i2, i1, i2[len(i2)-1], i1[0])
			currMax = i2[len(i2)-1] - i1[0]
			if currMax < 0 {
				currMax = -currMax
			}
			// m1 := i2[len(i2)-1] - i1[0]
			// m2 := i1[len(i1)-1] - i2[0]
			// if m1 > m2 {
			// 	currMax = m1
			// } else {
			// 	currMax = m2
			// }
		}
		// fmt.Println(currMax)
		if currMax > max {
			max = currMax
		}
	}
	for d, i1 := range ranges[1] {
		if i2, ok := ranges[1][d-1]; ok {
			// fmt.Println(i2, i1, i2[len(i2)-1], i1[0])
			currMax = i2[len(i2)-1] - i1[0]
			if currMax < 0 {
				currMax = -currMax
			}
			// m1 := i2[len(i2)-1] - i1[0]
			// m2 := i1[len(i1)-1] - i2[0]
			// if m1 > m2 {
			// 	currMax = m1
			// } else {
			// 	currMax = m2
			// }
		}
		// fmt.Println(currMax)
		if currMax > max {
			max = currMax
		}
		// d - 1
		if i2, ok := ranges[0][d-1]; ok {
			// fmt.Println(i2, i1, i2[len(i2)-1], i1[0])
			currMax = i2[len(i2)-1] - i1[0]
			if currMax < 0 {
				currMax = -currMax
			}
			// m1 := i2[len(i2)-1] - i1[0]
			// m2 := i1[len(i1)-1] - i2[0]
			// if m1 > m2 {
			// 	currMax = m1
			// } else {
			// 	currMax = m2
			// }
		}
		// fmt.Println(currMax)
		if currMax > max {
			max = currMax
		}
	}
	// fmt.Println(ranges[0])
	// fmt.Println(ranges[1])

	fmt.Print(max + 1)
}
