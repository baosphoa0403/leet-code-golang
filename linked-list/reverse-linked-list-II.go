package linkedlist

func ReverseBetween() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	// printList(head)
	out := reverseBetween(head, 2, 4)
	printList(out)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	start := prev.Next
	curr := start

	var reversePrev *ListNode = nil
	for i := 0; i < (right-left)+1; i++ {
		next := curr.Next
		curr.Next = reversePrev
		reversePrev = curr
		curr = next
	}

	prev.Next = reversePrev

	start.Next = curr

	return dummy.Next
}
