package twopointer

import "fmt"

func ReverseString() {
	inputs := [][]byte{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'a', 'b', 'c'}}

	for _, v := range inputs {
		fmt.Printf("before result: %v\n", v)
		execReverseString(v)
		fmt.Printf("after result: %v\n", v)
	}
}

func execReverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		fmt.Println("left", left, "- right", right)
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
