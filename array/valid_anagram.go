package array

import (
	"fmt"
	"strings"
)

func ValidAnagram() {
	tests := [][]string{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"aabb", "abb1"},
	}

	for _, tt := range tests {
		fmt.Printf("%s & %s → %v\n", tt[0], tt[1], isAnagram(tt[0], tt[1]))
	}
}

func isAnagram(s string, t string) bool {
	len1 := len(s)
	len2 := len(t)
	if len2 != len1 {
		return false
	}

	arr1 := strings.Split(s, "")
	arr2 := strings.Split(t, "")
	seen := make(map[string]int)

	for _, v := range arr1 {
		seen[v]++
	}

	for _, v := range arr2 {
		seen[v]--
		if seen[v] < 0 {
			return false
		}
	}

	return true
}
