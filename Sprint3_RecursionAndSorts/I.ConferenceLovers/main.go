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
	if n == 0 {
		return
	}
	s := bufio.NewScanner(os.Stdin)

	s.Scan()
	ids := strings.Split(strings.TrimSpace(s.Text()), " ")

	s.Scan()
	k, _ := strconv.Atoi(s.Text())
	cnt := make([]int, 10000)
	for _, id := range ids {
		i, _ := strconv.Atoi(id)
		cnt[i]++
	}

	cntOrdered := make([]int, 10000)
	for id, c := range cnt {
		cntOrdered[id] = c
	}

	sort.Ints(cntOrdered)

	// fmt.Println(cnt, cntOrdered)
	res := make([]int, k)
	for i := 0; i < k; i++ {
	LOOP:
		for j := len(cntOrdered) - 1; j >= 0; j-- {
			for m := 0; m < len(cnt); m++ {
				// fmt.Println(cnt[m], cntOrdered[j])
				if cnt[m] == cntOrdered[j] {
					res[i] = m
					cnt[m] = -1
					break LOOP
				}
			}
		}
	}

	for i, v := range res {
		if i == len(res)-1 {
			fmt.Printf("%d", v)
			continue
		}
		fmt.Printf("%d ", v)
	}
}
