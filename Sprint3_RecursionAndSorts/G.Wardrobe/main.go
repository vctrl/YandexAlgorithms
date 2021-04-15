package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	if n == 0 {
		fmt.Printf(" ")
		return
	}

	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 1000000*2+1000)

	s.Scan()

	arr := strings.Split(strings.TrimSpace(s.Text()), " ")
	// fmt.Println(arr)

	var i, j = 0, n - 1
	res := make([]string, n)

	for _, e := range arr {
		switch e {
		case "0":
			res[i] = "0"
			i++
		case "2":
			res[j] = "2"
			j--
		}
	}

	for k := i; k <= j; k++ {
		res[k] = "1"
	}

	// var i int
	// for j, cnt := range im {
	// 	for cnt > 0 {
	// 		arr[i] = strconv.Itoa(j)
	// 		i++
	// 		cnt--
	// 	}
	// }

	var result bytes.Buffer
	result.Grow(n*4 + n)
	for i := 0; i < n; i++ {
		result.WriteString(res[i])
		if i != n-1 {
			result.WriteString(" ")
		}
	}

	fmt.Print(result.String())
	// for i, v := range res {
	// 	if i == len(arr)-1 {
	// 		fmt.Printf("%s", v)
	// 	} else {
	// 		fmt.Printf("%s ", v)
	// 	}
	// }
}
