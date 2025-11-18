package twopointer

import (
	"fmt"
	"unicode"
)

// Level 1 – Basic Foundation (2 pointers opposite direction)
func IsPalindrome() {
	inputs := []string{"A man, a plan, a canal: Panama"}

	for _, v := range inputs {
		res := execIsPalindrome(v)

		fmt.Printf("input: %v  -> result: %t\n", v, res)
	}

}

func execIsPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !isAlnum(rune(s[left])) {
			left++
		}

		for left < right && !isAlnum(rune(s[right])) {
			right--
		}
		fmt.Printf("after left: %d - %c - right: %d - %c\n", left, rune(s[left]), right, rune(s[right]))

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		// fmt.Printf("left: %c - right: %c\n", rune(s[left]), rune(s[right]))
		left++
		right--
	}

	return true
}

func isAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
