package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"testing"
	"time"
)

const (
	tlen = 10
	slen = 100
)

func TestSubsequence(t1 *testing.T) {
	// var y int
	// for i := 0; i < 10000; i++ {
	// 	t, s := generateStrings()

	// 	if !CheckChar(0, 0, t, s) {
	// 		fmt.Println(t, s)
	// 	}
	// 	y++
	// }
	// fmt.Println(y)

	// test 2
	// for i := 0; i < 150000; i++ {
	// 	len := 150000
	// 	s := generateString(len)
	// 	rand.Seed(time.Now().UnixNano())
	// 	r1 := rand.Intn(len)
	// 	r2 := rand.Intn(len)
	// 	var t string
	// 	if r1 < r2 {
	// 		t = s[r1:r2]
	// 	} else {
	// 		t = s[r2:r1]
	// 	}
	// 	// fmt.Println(s)

	// 	fmt.Println(CheckChar(0, 0, t, s))
	// }

	// var sb strings.Builder
	var sb1 strings.Builder

	for i := 0; i < 150000; i++ {
		// sb.Write([]byte("s"))
		sb1.Write([]byte("s"))
	}

	f, err := os.Create("input1.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	// f.WriteString(sb.String())
	f.WriteString("a")
	f.WriteString("\n")
	f.WriteString(sb1.String())
}

func generateString(len int) string {
	min := int("a"[0])
	max := int("z"[0])
	var t string

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		r := rand.Intn(max-min) + min
		t = t + string(rune(r))
	}

	return t
}

func generateStrings() (string, string) {
	min := int("a"[0])
	max := int("z"[0])
	var t string
	indexes := make([]int, tlen)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < tlen; i++ {
		r := rand.Intn(max-min) + min
		t = t + string(rune(r))

		var rInd int
	LOOP:
		for {
			rInd = rand.Intn(slen)
			for _, i := range indexes {
				if i == rInd {
					continue LOOP
				}
			}

			break
		}

		indexes[i] = rInd
	}

	var sb strings.Builder
	sort.Ints(indexes)

	j := 0
	for i := 0; i < slen; i++ {
		var r rune
		if j < tlen && i == indexes[j] {
			r = rune(t[j])
			j++
		} else {
			r = rune(rand.Intn(max-min) + min)
		}

		sb.WriteRune(r)
	}

	return t, sb.String()
}
