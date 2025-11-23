package slidewindow

import (
	"fmt"
	"math"
)

func MinWindow() {
	inputs := map[string]string{
		"ADOBECODEBANC": "ABC",
		"a":             "a",
	}

	for k, v := range inputs {
		// fmt.Println(k, v)
		res := minWindow(k, v)
		fmt.Printf("input: %v  -> result: %v\n", k, res)
	}
}

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	window := make(map[byte]int)

	left := 0
	right := 0
	valid := 0

	start := 0
	minLen := math.MaxInt32

	required := len(need)

	for right < len(s) {
		c := s[right]
		right++

		if _, exists := need[c]; exists {
			window[c]++

			if window[c] == need[c] {
				valid++
			}
		}

		for valid == required {
			if right-left < minLen {
				minLen = right - left
				start = left
			}

			d := s[left]
			left++

			if _, exists := need[d]; exists {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[start : start+minLen]
}
