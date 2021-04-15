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
	Print(&node1)
	result := Solution(&node1, 2)
	fmt.Println("***after***")
	Print(result)
	fmt.Println("***initial***")
	Print(&node1)
}

func Print(head *ListNode) {
	if head == nil {
		return
	}

	fmt.Println(head.data)
	Print(head.next)
}

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		head = head.next
		return head
	}

	prev := head
	for idx > 1 {
		prev = prev.next
		idx--
	}

	prev.next = prev.next.next
	return head
}
