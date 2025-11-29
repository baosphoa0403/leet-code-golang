package linkedlist

func PartitionList() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}

	printList(head)
	out := partition(head, 3)
	printList(out)
}

func partition(head *ListNode, x int) *ListNode {
	smallDummy := &ListNode{}
	largeDummy := &ListNode{}

	small := smallDummy
	large := largeDummy

	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = nil

		if cur.Val < x {
			small.Next = cur
			small = small.Next
		} else {
			large.Next = cur
			large = large.Next
		}
		cur = next
	}

	small.Next = largeDummy.Next
	return smallDummy.Next
}
