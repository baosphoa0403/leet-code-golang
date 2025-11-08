package array

import "fmt"

func LongestConsecutive() {
	inputs := [][]int{{100, 4, 200, 1, 3, 2}, {0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, {1, 0, 1, 2}}

	for _, v := range inputs {
		res := execLongestConsecutive(v)
		fmt.Printf("input: %v  -> result: %d\n", v, res)
	}
}

func execLongestConsecutive(nums []int) int {
	// start := 0
	set := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		set[nums[i]] = true
	}

	maxLen := 0
	for n := range set {
		if !set[n-1] {

			length := 1
			current := n

			for set[current+1] {
				fmt.Printf("current: %d -> %t, next-current: %d -> %t\n", current, set[current], current+1, set[current+1])
				current++
				length++
			}

			if length > maxLen {
				maxLen = length
			}

		}
	}

	return maxLen
}
