package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Segment struct {
	begin int
	end   int
}

type ByBegin []Segment

func (b ByBegin) Len() int           { return len(b) }
func (b ByBegin) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByBegin) Less(i, j int) bool { return b[i].begin < b[j].begin }

func main() {
	var n int
	fmt.Scanf("%d", &n)
	s := bufio.NewScanner(os.Stdin)
	ss := make([]Segment, n)
	i := 0
	for s.Scan() {
		seg := strings.Split(s.Text(), " ")
		b, _ := strconv.Atoi(seg[0])
		e, _ := strconv.Atoi(seg[1])
		ss[i] = Segment{begin: b, end: e}
		i++
	}

	// sort

	sort.Sort(ByBegin(ss))

	// for _, s := range ss {
	// 	fmt.Printf("%d %d\n", s.begin, s.end)
	// }
	result := make([]*Segment, 0)
	minBegin, maxEnd := -1, -1
	for i := 0; i < n; i++ {
		if minBegin < 0 {
			minBegin = ss[i].begin
		}
		if maxEnd < 0 {
			maxEnd = ss[i].end

			if n == 1 {
				result = append(result, &Segment{begin: minBegin, end: maxEnd})
			}

			continue
		}

		if maxEnd < ss[i].begin {
			result = append(result, &Segment{begin: minBegin, end: maxEnd})
			minBegin, maxEnd = ss[i].begin, ss[i].end

			if i == n-1 {
				result = append(result, &Segment{begin: minBegin, end: maxEnd})
			}
		} else {
			if ss[i].end > maxEnd {
				maxEnd = ss[i].end
			}

			if i == n-1 {
				result = append(result, &Segment{begin: minBegin, end: maxEnd})
			}
		}
	}

	// fmt.Println("result")
	for _, s := range result {
		fmt.Printf("%d %d\n", s.begin, s.end)
	}
}
