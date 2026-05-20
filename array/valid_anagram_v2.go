package array

import (
	"fmt"
	"strings"
)

func IsAnagramV2() bool {
	s := "anagram"
	t := "nagara"

	len1 := len(s)
	len2 := len(t)
	if len2 != len1 {
		return false
	}

	arr1 := strings.Split(s, "")
	arr2 := strings.Split(t, "")
	seen := map[string]int{}
	for _, v := range arr1 {
		seen[v]++
	}

	for _, v := range arr2 {
		seen[v]--
		fmt.Println("seen[v]: ", seen[v])
		if val := seen[v]; val < 0 {
			return false
		}
	}

	return true
}
