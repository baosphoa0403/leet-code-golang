package linkedlist

import "fmt"

func DetectCycle() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	//n3 := &ListNode{Val: 3}
	//n4 := &ListNode{Val: 4}
	//n5 := &ListNode{Val: 5}

	n1.Next = n2
	n2.Next = n1
	//n3.Next = n4
	//n4.Next = n5
	//n5.Next = n3 // tạo cycle: 5 -> 3

	out := detectCycle(n1)
	println(out.Val)
}
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		fmt.Println("slow", slow, "fast", fast)
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	//fmt.Println("before slow", slow)
	slow = head
	fmt.Println("after slow", slow, fast)

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow // start of cycle
}
