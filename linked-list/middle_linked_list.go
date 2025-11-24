package linkedlist

import "fmt"

func MiddleNode() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	out := middleNode(head)
	printList(out)
	fmt.Println(out.Val)
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	// fmt.Println("slow", slow, fast)

	for fast.Next != nil {
		fmt.Println("slow", slow, fast)
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
