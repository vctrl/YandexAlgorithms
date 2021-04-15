package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var l int

	fmt.Scanf("%d\n", &l)

	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 1000000*1000000+1000000)

	s.Scan()

	numsStr := strings.Split(s.Text(), " ")

	nums := make([]int, l)
	for i := 0; i < l; i++ {
		n, _ := strconv.Atoi(numsStr[i])
		nums[i] = n
	}

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	res := make([]int, 2)
	res[0] = binarySearch(nums, 0, len(nums), n)

	if res[0] == 0 || res[0] == len(nums)+1 {
		res[1] = -1
		res[0] = -1
	} else {
		res[1] = binarySearch(nums, 0, len(nums), n*2)
		if res[1] == 0 || res[1] == len(nums)+1 {
			res[1] = -1
		}
	}

	for _, r := range res {
		fmt.Printf("%d ", r)
	}
}

func binarySearch(arr []int, i, j, n int) int {
	if i >= j {
		return i + 1
	}

	mid := (i + j) / 2
	if arr[mid] >= n {
		return binarySearch(arr, i, mid, n)
	}

	return binarySearch(arr, mid+1, j, n)
}
