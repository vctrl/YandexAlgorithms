package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cmdCount int
	fmt.Scanf("%d", &cmdCount)
	s := bufio.NewScanner(os.Stdin)
	q := NewLinkedQueue()
	for s.Scan() {
		parseCommand(q, s.Text())
	}
}

type Node struct {
	value int
	Next  *Node
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

type LinkedQueue struct {
	head *Node
	tail *Node
	size int
}

func (q *LinkedQueue) Get() (int, bool) {
	if q.size == 0 {
		return 0, false
	}

	res := q.head.value
	q.head = q.head.Next
	q.size--
	return res, true
}

func (q *LinkedQueue) Put(x int) {
	node := &Node{value: x}
	if q.size == 0 {
		q.tail = node
		q.head = node
		q.size++
		return
	}

	q.tail.Next = node
	q.tail = node
	q.size++
}

func (q *LinkedQueue) Size() int {
	return q.size
}

func parseCommand(q *LinkedQueue, cmd string) {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "get":
		x, ok := q.Get()
		if !ok {
			fmt.Println("error")
		} else {
			fmt.Println(x)
		}
	case "put":
		x, _ := strconv.Atoi(parts[1])
		q.Put(x)
	case "size":
		fmt.Println(q.Size())
	}
}
