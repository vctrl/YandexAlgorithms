package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	q     = 1000
	R     = 123987123
	chars = "qwertyuiopasdfghjklzxcvbnm"
)

func main() {
	// fmt.Println(hash("grocyo"))
	// fmt.Println(hash("ndpjjb"))
	rand.Seed(time.Now().UnixNano())

	hs := make(map[int]string)
	for {
		s := getRandomStr(6)
		h := hash(s)
		fmt.Println(s, h)
		if ss, ok := hs[h]; ok && ss != s {
			fmt.Printf("success\n%s\n%s\n", s, hs[h])
			break
		} else {
			hs[h] = s
		}
	}
}

func getRandomStr(l int) string {
	b := make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = byte(chars[rand.Intn(len(chars)-1)])
	}

	return string(b)
}

func hash(str string) int {
	var h int
	li := len(str) - 1
	for _, s := range str[:li] {
		h = (h + int(s)) * q

		if h > R {
			h %= R
		}
	}

	h += int(str[li])
	return h % R
}
