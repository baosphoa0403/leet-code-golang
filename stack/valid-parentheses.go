package stack

import "fmt"

func ValidParentheses() {
	s := "()[]{}"
	out := isValid(s)
	fmt.Println(out)
}

func isValid(s string) bool {
	stack := []rune{}

	open := map[rune]bool{'(': true, '[': true, '{': true}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		if open[v] {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != pairs[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
