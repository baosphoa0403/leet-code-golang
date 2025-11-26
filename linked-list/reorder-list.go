package linkedlist

import "fmt"

func ReorderList() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	reorderReverse(head)
	//printList(head)
	//reorderList(head)
	//printList(head)
}

func reorderList(head *ListNode) {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fmt.Println("slow", slow, "fast", fast)
		slow = slow.Next
		fast = fast.Next.Next
		fmt.Println("slow", slow, "fast", fast)
	}
	middle := slow
	fmt.Println("mid", middle.Next)
	printList(middle)
}

func reorderReverse(head *ListNode) {
	printList(head)
	prev := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
		fmt.Println("prev", prev, head)
	}

	printList(prev.Next)
}
