package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
}

func main() {
	node3 := ListNode{data: "mazafak3", next: nil}
	node2 := ListNode{data: "mazafak2", next: &node3}
	node1 := ListNode{data: "mazafak1", next: &node2}
	// Print(&node1)
	fmt.Println(Solution(&node1, "mazafak228"))
}

func Print(head *ListNode) {
	if head == nil {
		return
	}

	fmt.Println(head.data)
	Print(head.next)
}

func Solution(head *ListNode, elem string) int {
	i := 0
	for head != nil {
		if head.data == elem {
			return i
		}

		head = head.next
		i++
	}

	return -1
}
