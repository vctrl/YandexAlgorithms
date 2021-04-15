package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 100*1000+100)
	s.Scan()
	str := strings.Split(s.Text(), " ")

	sort.Sort(byCategory(str))

	// fmt.Print(strings.Join(str, " "))
	fmt.Print(strings.Join(str, ""))
}

func Sort(str []string) string {
	sort.Sort(byCategory(str))
	return strings.Join(str, "")
}

type byCategory []string

func (s byCategory) Less(i, j int) bool {
	var min, max int
	if len(s[i]) < len(s[j]) {
		min = len(s[i])
		max = len(s[j])
	} else {
		min = len(s[j])
		max = len(s[i])
	}

	for k := 0; k < min; k++ {
		if s[i][k] < s[j][k] {
			return false
		} else if s[i][k] > s[j][k] {
			return true
		}
	}

	if len(s[i]) == len(s[j]) {
		return false
	}

	if min == len(s[j]) {
		for k := 0; k < max; k++ {
			if s[i][k] < s[i][(k+min)%max] {
				return true
			} else if s[i][k] > s[i][(k+min)%max] {
				return false
			}
		}

		return true
	}

	for k := 0; k < max; k++ {
		if s[j][k] < s[j][(k+min)%max] {
			return false
		} else if s[j][k] > s[j][(k+min)%max] {
			return true
		}
	}

	return false
}

func (s byCategory) Len() int {
	return len(s)
}

func (s byCategory) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
