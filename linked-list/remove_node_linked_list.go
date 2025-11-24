package linkedlist

func RemoveNthFromEnd() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	// head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	// head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}

	printList(head)
	out := removeNthFromEnd(head, 2)
	printList(out)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
