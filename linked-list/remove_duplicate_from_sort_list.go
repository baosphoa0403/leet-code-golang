package linkedlist

func RemoveDuplicateFromSortList() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 5}

	printList(head)
	out := deleteDuplicates(head)
	printList(out)
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			val := curr.Val
			for curr != nil && curr.Val == val {
				curr = curr.Next
			}
			prev.Next = curr
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}
