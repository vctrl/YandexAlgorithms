package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cmdCount, cap int
	fmt.Scanf("%d\n%d", &cmdCount, &cap)
	s := bufio.NewScanner(os.Stdin)
	q := NewSizedQueue(cap)
	// var result string
	for s.Scan() {
		parseCommand(q, s.Text())
	}
}

type SizedQueue struct {
	queue []int
	head  int
	tail  int
	cap   int
	size  int
}

func NewSizedQueue(size int) *SizedQueue {
	return &SizedQueue{queue: make([]int, size), head: 0, tail: 0, size: 0, cap: size}
}

func (s *SizedQueue) Push(x int) bool {
	if s.size == s.cap {
		return false
	}

	s.queue[s.tail] = x
	s.size++

	s.tail = (s.tail + 1) % s.cap
	return true
}

func (s *SizedQueue) Pop() (int, bool) {
	if s.size == 0 {
		return 0, false
	}

	result := s.queue[s.head]
	s.head = (s.head + 1) % s.cap
	s.size--
	return result, true
}

func (s *SizedQueue) Peek() (int, bool) {
	if s.size == 0 {
		return 0, false
	}

	return s.queue[s.head], true
}

func (s *SizedQueue) Size() int {
	return s.size
}

func parseCommand(q *SizedQueue, cmd string) {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "push":
		x, _ := strconv.Atoi(parts[1])
		ok := q.Push(x)
		if !ok {
			fmt.Println("error")
		}
	case "pop":
		res, ok := q.Pop()
		if !ok {
			fmt.Println("None")
		} else {
			fmt.Println(res)
		}
	case "peek":
		res, ok := q.Peek()
		if !ok {
			fmt.Println("None")
		} else {
			fmt.Println(res)
		}
	case "size":
		fmt.Println(q.Size())
	}
}
