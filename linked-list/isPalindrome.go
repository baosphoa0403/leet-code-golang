package linkedlist

import "fmt"

func IsPalindrome() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}

	printList(head)
	res := isPalindrome(head)
	fmt.Println("res:", res)
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		fmt.Println("slow", slow, "fast", fast)
	}

	printList(head)

	fmt.Println("final", slow)

	return false
}
