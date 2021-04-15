package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	scanner.Buffer(buf, 150000*2+10000)
	scanner.Scan()
	t := scanner.Text()
	if t == "" {
		fmt.Print("True")
		return
	}
	scanner.Scan()
	s := scanner.Text()

	res := CheckChar(0, 0, t, s)
	var strRes string

	if res {
		strRes = "True"
	} else {
		strRes = "False"
	}

	fmt.Println(strRes)
}

func CheckChar(ti, si int, t, s string) bool {
	if ti == len(t) || si == len(s) {
		return false
	}

	if s[si] == t[ti] {
		if ti == len(t)-1 {
			return true
		}

		return CheckChar(ti+1, si+1, t, s)
	}

	return CheckChar(ti, si+1, t, s)
}
