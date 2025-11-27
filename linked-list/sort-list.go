package linkedlist

import (
	"fmt"
	"math"
)

func SortList() {
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 3}

	printList(head)
	out := sortList(head)
	printList(out)
}
func sortList(head *ListNode) *ListNode {
	slow, fast := head, head
	dummy := &ListNode{Val: math.MinInt}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fmt.Println("slow:", slow)

	return dummy.Next
}
