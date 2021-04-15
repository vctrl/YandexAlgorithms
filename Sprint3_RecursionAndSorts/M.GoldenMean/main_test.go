package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestMedian(t *testing.T) {
	// for c := 0; c < 100; c++ {
	// r := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(r)

	a1 := []int{4, 4, 5, 7, 7, 7, 8, 9, 9, 10, 11}
	a2 := []int{0, 0, 0, 1, 3}
	// a1 := make([]int, 10000)
	// for i := 0; i < len(a1); i++ {
	// 	a1[i] = r1.Intn(10000)
	// }

	// a2 := make([]int, 10000)
	// for i := 0; i < len(a2); i++ {
	// 	a2[i] = r1.Intn(10000)
	// }

	sort.Ints(a1)
	sort.Ints(a2)
	// fmt.Println(a1, a2)

	a1a2 := merge(a1, a2)
	sort.Ints(a1a2)
	expected := getMedian(a1a2)

	fact := findMedian(a1, a2, len(a1), len(a2))

	if expected != fact {
		fmt.Printf("error: expected %v but was %v\n", expected, fact)
	} else {
		fmt.Printf("success %v %v\n", expected, fact)
	}
	//  else {
	// 	fmt.Printf("success %v %v", expected, fact)
	// }
	// fmt.Println(expected)
	// fmt.Println(a1a2)
	// }
}

func merge(a1, a2 []int) []int {
	res := make([]int, len(a1)+len(a2))
	for i := 0; i < len(a1); i++ {
		res[i] = a1[i]
	}
	for i := 0; i < len(a2); i++ {
		res[len(a1)+i] = a2[i]
	}

	return res
}

func getMedian(arr []int) float64 {
	if len(arr)%2 == 1 {
		return float64(arr[len(arr)/2])
	}

	return float64(arr[len(arr)/2]+arr[len(arr)/2-1]) / 2

}
