package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		rand.Seed(time.Now().UnixNano())
		r1 := rand.Intn(1000)
		r2 := rand.Intn(1000)

		r1str := strconv.Itoa(r1)
		r2str := strconv.Itoa(r2)

		expected1 := r1str + r2str
		exp1num, _ := strconv.Atoi(expected1)
		expected2 := r2str + r1str
		exp2num, _ := strconv.Atoi(expected2)

		var expected int
		if exp1num >= exp2num {
			expected = exp1num
		} else {
			expected = exp2num
		}

		fact := Sort([]string{r1str, r2str})
		factNum, _ := strconv.Atoi(fact)

		// fmt.Printf("test case %d: values %d %d expected %d fact %d\n", i, r1, r2, expected, factNum)
		if factNum != expected {
			t.Errorf("Values doesn't match: %d %d\n args:%d %d", factNum, expected, r1, r2)
		}
	}
}
