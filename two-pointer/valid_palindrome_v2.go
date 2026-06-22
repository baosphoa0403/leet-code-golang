package twopointer

import "strings"

func IsPalindromeV2(s string) bool {

	chars := []rune(s)

	left := 0
	right := len(chars) - 1
	for left < right {
		for left < right && !isAlphaNumeric(chars[left]) {
			left++
		}

		for left < right && !isAlphaNumeric(chars[right]) {
			right--
		}
		abc := string(chars[left])
		def := string(chars[right])

		if !strings.EqualFold(abc, def) {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphaNumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}
