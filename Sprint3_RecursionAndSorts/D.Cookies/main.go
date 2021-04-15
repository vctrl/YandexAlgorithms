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
	var nc, mc int
	s := bufio.NewScanner(os.Stdin)

	s.Scan()
	nc, _ = strconv.Atoi(s.Text())
	if nc == 0 {
		fmt.Print("0")
		return
	}

	s.Scan()
	nsStr := strings.Split(strings.TrimSpace(s.Text()), " ")
	ns := make([]int, nc)
	for i, n := range nsStr {
		ns[i], _ = strconv.Atoi(n)
	}

	s.Scan()
	mc, _ = strconv.Atoi(s.Text())
	if mc == 0 {
		fmt.Println("0")
		return
	}

	s.Scan()
	msStr := strings.Split(strings.TrimSpace(s.Text()), " ")
	ms := make([]int, mc)
	for i, m := range msStr {
		ms[i], _ = strconv.Atoi(m)
	}

	sort.Ints(ns)
	sort.Ints(ms)

	// fmt.Println(ns, ms)

	k := 0
	for i := 0; i < len(ms); i++ {
		// fmt.Println(i, ns[k], ms[i])
		if k >= len(ns) {
			break
		}

		if ns[k] <= ms[i] {
			k++
		}
	}

	fmt.Print(k)
}
