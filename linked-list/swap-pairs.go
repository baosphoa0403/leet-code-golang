package linkedlist

func SwapPairs() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	printList(head)
	out := swapPairs(head)
	printList(out)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy

	curr := head
	for curr != nil && curr.Next != nil {
		first := curr
		second := curr.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
		curr = first.Next
	}

	return dummy.Next
}
