package linkedlist

func InsertionSortList() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next.Next = &ListNode{Val: 8}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	out := insertionSortList(head)
	printList(out)

}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := head

	for curr != nil {
		next := curr.Next
		prev := dummy
		for prev.Next != nil && prev.Next.Val < curr.Val {
			prev = prev.Next
		}

		curr.Next = prev.Next
		prev.Next = curr
		curr = next
	}

	return dummy
}
