package array

import "fmt"

func MissingNumber() {
	inputs := [][]int{{3, 0, 1}, {0, 1}, {9, 6, 4, 2, 3, 5, 7, 0, 1}}

	for _, v := range inputs {
		// res := execMissingNumber(v)
		res := missingNumberV2(v)

		fmt.Printf("input: %v  -> result: %d\n", v, res)
	}
}

func execMissingNumber(nums []int) int {
	n := len(nums)
	set := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		set[nums[i]] = true
	}

	start := 0
	for start <= n {
		if !set[start] {
			return start
		}
		start++
	}

	return -1
}

func missingNumberV2(nums []int) int {
	xor := 0
	n := len(nums)
	for i := 0; i <= n; i++ {
		xor ^= i
	}

	for _, num := range nums {
		xor ^= num
	}

	return xor
}
