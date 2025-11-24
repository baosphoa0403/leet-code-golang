package linkedlist

import "fmt"

func IsPalindrome() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}

	res := isPalindrome(head)
	fmt.Println("res:", res)
}

func Reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversed := Reverse(slow)
	first := head
	second := reversed

	for first != nil && second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}

	return true
}
