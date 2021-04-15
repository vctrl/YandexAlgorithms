package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

func main() {
	node3 := ListNode{data: "mazafak3"}
	node2 := ListNode{data: "mazafak2"}
	node1 := ListNode{data: "mazafak1"}
	node1.next = &node2
	node2.prev = &node1
	node2.next = &node3
	node3.prev = &node2
	Print(&node1)
	res := Solution(&node1)
	Print(res)
}

func Print(head *ListNode) {
	if head == nil {
		return
	}

	fmt.Println(head.data)
	Print(head.next)
}

func Solution(head *ListNode) *ListNode {
	for {
		t := head.prev
		head.prev = head.next
		head.next = t

		if head.prev == nil {
			return head
		}

		head = head.prev
	}
}
