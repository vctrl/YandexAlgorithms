/*

ID посылки 49720106

-- ПРИНЦИП РАБОТЫ --

Сначала ищем индекс, в котором прерывается рост, используя бинарный поиск
Потом делаем бинарный поиск, учитывая сдвиг, высчитанный на шаге 1

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
По сути решение представляет чуть модифицированный бинарный поиск :)


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

O(logN+logN) = O(2logN) = O(logN)

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanf("%d\n%d\n", &n, &k)

	s := bufio.NewScanner(os.Stdin)

	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 10e3*5+10e3)
	s.Scan()

	str := strings.Split(s.Text(), " ")
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(str[i])
	}

	if n == 1 && nums[0] == k {
		fmt.Println(0)
		return
	}

	threshold := threshholdBSearch(nums, k, 0, n-1, 0, n-1)
	res := bSearchInBrokenArr(nums, k, 0, n-1, threshold, n)
	fmt.Println(res)
}

func threshholdBSearch(arr []int, n, l, r, le, re int) int {
	if arr[re] > arr[le] {
		return 0
	}

	mid := l + ((r - l) / 2)

	if l == r {
		return -1
	}

	if arr[mid] > arr[mid+1] {
		return mid + 1
	}

	if arr[mid] > arr[le] {
		return threshholdBSearch(arr, n, mid+1, r, le, re)
	}

	return threshholdBSearch(arr, n, l, mid, le, re)
}

// t for threshold
func bSearchInBrokenArr(arr []int, n, l, r, t, mod int) int {
	mid := l + ((r - l) / 2)
	midShifted := (mid + t) % mod
	midVal := arr[midShifted]

	if midVal == n {
		return midShifted
	}

	if r-l == 0 {
		return -1
	}

	if n > midVal {
		return bSearchInBrokenArr(arr, n, mid+1, r, t, mod)
	}

	return bSearchInBrokenArr(arr, n, l, mid, t, mod)
}
