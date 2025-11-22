package slidewindow

import "fmt"

func LengthOfLongestSubstring() {
	inputs := []string{
		// "abcabcbb",
		"pwwkew",
	}

	for _, k := range inputs {
		res := lengthOfLongestSubstring(k)
		fmt.Printf("input: %v  -> result: %v\n", k, res)
	}
}

func lengthOfLongestSubstring(s string) int {
	freq := map[byte]int{}
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		freq[c]++

		fmt.Println("freq", freq)
		for freq[c] > 1 {
			freq[s[left]]--
			left++
		}

		windowLength := right - left + 1
		if windowLength > maxLen {
			maxLen = windowLength
		}
	}

	return maxLen
}
