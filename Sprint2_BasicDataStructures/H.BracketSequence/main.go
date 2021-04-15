package main

import (
	"bufio"
	"fmt"
	"os"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type Stack struct {
	slice []rune
}

func (s *Stack) Init() {
	s.slice = make([]rune, 0)
}

func (s *Stack) Push(x rune) {
	// fmt.Println("pushing", x, "into stack")
	s.slice = append(s.slice, x)
}

func (s *Stack) Pop() (rune, error) {
	if len(s.slice) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	n := len(s.slice) - 1
	x := s.slice[n]
	// fmt.Println("popping", x, "from stack")
	s.slice = s.slice[:n]

	return x, nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.slice) == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	stack := &Stack{}
	stack.Init()

	scanner.Scan()
	// [], (), {}.
	bracketSequence := scanner.Text()

	result := true
	for _, c := range bracketSequence {
		if isOpening(c) {
			stack.Push(c)
		} else {
			lastBracket, _ := stack.Pop()
			if getPair(lastBracket) != string(c) {
				result = false
				break
			}
		}
	}

	if result && stack.IsEmpty() {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func getPair(c rune) string {
	if string(c) == "[" {
		return "]"
	} else if string(c) == "{" {
		return "}"
	} else {
		return ")"
	}
}

func isOpening(c rune) bool {
	return string(c) == "[" || string(c) == "{" || string(c) == "("
}

func isClosing(c rune) bool {
	return string(c) == "]" || string(c) == "}" || string(c) == ")"
}
