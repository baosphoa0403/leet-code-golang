package linkedlist

import "fmt"

func MergeTwoLists() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 4}

	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 3}
	head1.Next.Next = &ListNode{Val: 4}

	out := mergeTwoLists(head, head1)
	printList(out)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	cur1 := list1
	cur2 := list2
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			tail.Next = cur1
			cur1 = cur1.Next
		} else {
			tail.Next = cur2
			cur2 = cur2.Next
		}
		tail = tail.Next
		fmt.Println(cur1, cur2)
	}

	if cur1 != nil {
		tail.Next = cur1
	}

	if cur2 != nil {
		tail.Next = cur2
	}

	return dummy.Next
}
