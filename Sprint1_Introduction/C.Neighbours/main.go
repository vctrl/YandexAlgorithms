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
	var rowCount, colCount int
	fmt.Scanf("%d\n%d\n", &rowCount, &colCount)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	lines := make([]string, rowCount)
	scanner.Split(bufio.ScanLines)
	for i := 0; i < rowCount; i++ {
		scanner.Scan()
		lines[i] = scanner.Text()
	}

	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	result := make([]int, 0)
	if y-1 >= 0 {
		numbers := strings.Split(lines[y-1], " ")
		up, _ := strconv.Atoi(numbers[x])
		result = append(result, up)
	}
	var numberslr []string
	if x-1 >= 0 {
		numberslr = strings.Split(lines[y], " ")
		left, _ := strconv.Atoi(numberslr[x-1])
		result = append(result, left)
	}
	if x+1 < colCount {
		if numberslr == nil {
			numberslr = strings.Split(lines[y], " ")
		}
		right, _ := strconv.Atoi(numberslr[x+1])
		result = append(result, right)
	}
	if y+1 < rowCount {
		numbers := strings.Split(lines[y+1], " ")
		down, _ := strconv.Atoi(numbers[x])
		result = append(result, down)
	}

	sort.Ints(result)
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
		if i != len(result)-1 {
			fmt.Print(" ")
		}
	}
}
