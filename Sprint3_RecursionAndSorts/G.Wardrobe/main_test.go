package main

import (
	"os"
	"strings"
	"testing"
)

func TestWardrobe(t *testing.T) {

	var sb strings.Builder

	for i := 0; i < 100000; i++ {
		sb.WriteString("0 ")
	}

	f, _ := os.Create("input1.txt")
	defer f.Close()

	f.WriteString(sb.String())
}
