package linkedlist

import "fmt"

func MiddleNode() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}

	printList(head)
	out := middleNode(head)
	printList(out)
	fmt.Println(out.Val)
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil {
		slow = slow.Next
		if fast.Next.Next == nil {
			break
		}
		fast = fast.Next.Next
	}

	return slow
}
