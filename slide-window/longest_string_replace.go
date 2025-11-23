package slidewindow

import "fmt"

func CharacterReplacement() {
	inputs := map[int]string{
		2: "ABAB",
		1: "AABABBA",
	}

	for v, k := range inputs {
		res := characterReplacement(k, v)
		fmt.Printf("input: %v  -> result: %v\n", k, res)
	}
}

func characterReplacement(s string, k int) int {
	freq := map[byte]int{}
	left := 0
	maxFreq := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		freq[c]++

		if freq[c] > maxFreq {
			maxFreq = freq[c]
		}

		windowSize := right - left + 1
		if windowSize-maxFreq > k {
			freq[s[left]]--
			left++
		} else {
			if windowSize > maxLen {
				maxLen = windowSize
			}
		}
	}

	return maxLen
}
