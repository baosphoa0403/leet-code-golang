package linkedlist

import "fmt"

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
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil

	left := sortList(head)
	right := sortList(mid)

	out := merge(left, right)
	return out
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			tail.Next = left
			left = left.Next
		} else {
			tail.Next = right
			right = right.Next
		}
		tail = tail.Next
	}

	if left != nil {
		tail.Next = left
	}
	if right != nil {
		tail.Next = right
	}
	fmt.Println("tail", tail, "left", left, "right", right)

	printList(tail)

	return dummy.Next
}
