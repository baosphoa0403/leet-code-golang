package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	printList(head)
	out := reverseList(head)
	printList(out)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		fmt.Printf("prev=%v curr=%v next=%v\n",
			val(prev), val(cur), val(next))
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func val(n *ListNode) int {
	if n == nil {
		return -1
	}
	return n.Val
}
