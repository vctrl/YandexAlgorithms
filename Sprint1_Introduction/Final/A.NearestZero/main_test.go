package main

import (
	"os"
	"testing"
	"time"
)

func BenchmarkAlgorithm(b *testing.B) {
	f, _ := os.Open("input.txt")

	defer f.Close()

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = f

	for i := 0; i < 10; i++ {
		main()
		time.Sleep(1 * time.Second)
	}

	// for i := 0; i < 1; i++ {
	// 	f, _ := os.Open("input.txt")
	// 	// s := bufio.NewScanner(f)
	// 	// s.Scan()
	// 	// fmt.Print(s.Text())

	// 	defer f.Close()
	// 	NearestZeros(f)
	// 	time.Sleep(1 * time.Second)
	// }
}
