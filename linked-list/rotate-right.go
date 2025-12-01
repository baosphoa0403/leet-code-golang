package linkedlist

func RotateRight() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	out := rotateRight(head, 2)
	printList(out)
}

func countNode(head *ListNode) (int, *ListNode) {
	count := 1
	tail := head
	for tail.Next != nil {
		count++
		tail = tail.Next
	}
	return count, tail
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	count, tail := countNode(head)
	k = k % count
	if k == 0 {
		return head
	}

	tail.Next = head

	steps := count - k - 1
	cur := head
	for steps > 0 {
		cur = cur.Next
		steps--
	}

	newHead := cur.Next
	cur.Next = nil

	return newHead
}
