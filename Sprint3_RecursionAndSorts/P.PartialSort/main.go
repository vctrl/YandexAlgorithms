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
	s.Scan()
	str := strings.Split(s.Text(), " ")

	parts := 0

	first := 0
	for i := 0; i < n; i++ {

		aux := make([]int, i+1)
		for j := first; j <= i; j++ {
			nn, _ := strconv.Atoi(str[j])
			// fmt.Println("aux", j, nn)
			aux[j] = nn
		}

		// fmt.Println(aux)
		sort.Ints(aux)

		isPart := true
		m := first
		for i := first; i < len(aux); i++ {
			// fmt.Println("checking", m, aux[i])
			if aux[i] != m {
				isPart = false
				break
			}
			m++
		}

		// fmt.Println("part?", isPart)
		if isPart {
			parts++
			first = i + 1
		}
	}

	fmt.Print(parts)
}
