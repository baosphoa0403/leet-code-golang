package twopointer

import (
	"fmt"
	"strings"
)

func ReverseWords() {
	inputs := []string{"Let's take LeetCode contest", "Mr Ding"}

	for _, v := range inputs {
		res := ExecReverseWords(v)
		fmt.Printf("input: %s  -> result: %s\n", v, res)
	}
}

func ExecReverseWords(s string) string {
	arr := strings.Split(s, " ")

	for index, v := range arr {
		s := []rune(v)
		left, right := 0, len(s)-1
		for left < right {
			s[left], s[right] = s[right], s[left]
			left++
			right--
		}
		arr[index] = string(s)
	}

	return strings.Join(arr, " ")
}
