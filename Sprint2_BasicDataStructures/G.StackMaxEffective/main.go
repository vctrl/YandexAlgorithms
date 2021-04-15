package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type Stack struct {
	slice []int

	maxValues []int
}

func (s *Stack) Init() {
	s.slice = make([]int, 0)
	s.maxValues = make([]int, 0)
}

func (s *Stack) Push(x int) {
	// fmt.Println("pushing", x, "into stack")
	s.slice = append(s.slice, x)

	// max value
	if len(s.maxValues) == 0 {
		s.maxValues = append(s.maxValues, x)
	} else if x >= s.maxValues[len(s.maxValues)-1] {
		s.maxValues = append(s.maxValues, x)
	}
}

func (s *Stack) Pop() (int, error) {
	if len(s.slice) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	n := len(s.slice) - 1
	x := s.slice[n]
	// fmt.Println("popping", x, "from stack")
	s.slice = s.slice[:n]

	// max value
	nmax := len(s.maxValues) - 1
	if len(s.maxValues) > 0 && x == s.maxValues[nmax] {
		s.maxValues = s.maxValues[:nmax]
	}

	return x, nil
}

func (s *Stack) GetMax() (int, error) {
	if len(s.slice) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	nmax := len(s.maxValues) - 1

	// fmt.Println("returning max value:", max)
	return s.maxValues[nmax], nil
}

func main() {
	var cnt string
	fmt.Scanf("%d", &cnt)

	scanner := bufio.NewScanner(os.Stdin)

	stack := &Stack{}
	stack.Init()

	for scanner.Scan() {
		cmd := scanner.Text()
		parseCmd(stack, cmd)
	}
}

func parseCmd(stack *Stack, rawCmd string) {
	// fmt.Println("--->executing", rawCmd)
	args := strings.Split(rawCmd, " ")
	switch args[0] {
	case "push":
		val, _ := strconv.Atoi(args[1])
		stack.Push(val)
	case "pop":
		_, err := stack.Pop()
		if err != nil {
			fmt.Println("error")
			// } else {
			// 	fmt.Println(val)
		}
	case "get_max":
		max, err := stack.GetMax()
		if err != nil {
			fmt.Println("None")
		} else {
			fmt.Println(max)
		}
	}
}
