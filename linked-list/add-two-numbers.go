package linkedlist

func AddTwoNumbers() {
	head := &ListNode{Val: 2}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}

	head2 := &ListNode{Val: 5}
	head2.Next = &ListNode{Val: 6}
	head2.Next.Next = &ListNode{Val: 4}
	// head := &ListNode{Val: 9}
	// head.Next = &ListNode{Val: 9}
	// head.Next.Next = &ListNode{Val: 9}
	// head.Next.Next.Next = &ListNode{Val: 9}
	// head.Next.Next.Next.Next = &ListNode{Val: 9}
	// head.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	// head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}

	// head2 := &ListNode{Val: 9}
	// head2.Next = &ListNode{Val: 9}
	// head2.Next.Next = &ListNode{Val: 9}
	// head2.Next.Next.Next = &ListNode{Val: 9}

	out := addTwoNumbers(head, head2)
	printList(out)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		digit := sum % 10
		carry = sum / 10

		cur.Next = &ListNode{Val: digit}
		cur = cur.Next
	}
	return dummy.Next
}
