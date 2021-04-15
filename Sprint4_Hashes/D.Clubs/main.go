package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 10000)

	clubs := make(map[string]struct{})
	var res strings.Builder

	for s.Scan() && n > 0 {
		n--
		t := s.Text()
		if _, ok := clubs[t]; !ok {
			res.WriteString(s.Text() + "\n")
			clubs[t] = struct{}{}
		}
	}

	fmt.Print(res.String())
}
