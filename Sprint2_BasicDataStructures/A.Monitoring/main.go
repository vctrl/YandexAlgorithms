package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var rowCnt, colCnt int
	fmt.Scan(&rowCnt)
	fmt.Scan(&colCnt)

	result := make([][]string, colCnt)
	for i := 0; i < len(result); i++ {
		result[i] = make([]string, rowCnt)
	}

	s := bufio.NewScanner(os.Stdin)
	for i := 0; i < rowCnt; i++ {
		s.Scan()
		row := strings.Split(s.Text(), " ")
		for j := 0; j < colCnt; j++ {
			result[j][i] = row[j]
		}
	}

	for i := 0; i < colCnt; i++ {
		fmt.Println(strings.Join(result[i], " "))
	}
}
