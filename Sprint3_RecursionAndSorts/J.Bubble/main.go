package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 1000*1000+1000)

	s.Scan()
	str := strings.Split(s.Text(), " ")
	arr := make([]int, n)
	for i := 0; i < len(str); i++ {
		arr[i], _ = strconv.Atoi(str[i])
	}

	for i := len(arr); i > 0; i-- {
		var swapped bool
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				t := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = t
				swapped = true
			}
		}

		if swapped {
			print(arr)
		} else if !swapped && i == len(arr) {
			print(arr)
			break
		}
	}
}

func print(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}
