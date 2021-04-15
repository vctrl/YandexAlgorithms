package main

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	s1 := &Student{Name: "gosha", Solved: 2, Penalty: 90}
	s2 := &Student{Name: "rita", Solved: 2, Penalty: 90}
	fmt.Println(s1.Compare(s2))
}
