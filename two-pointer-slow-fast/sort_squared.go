package twopointerslowfast

import (
	"fmt"
)

func SortedSquares() {
	inputs := [][]int{
		{-4, -1, 0, 3, 10},
		{-7, -3, 2, 3, 11},
	}

	for _, v := range inputs {
		input := append([]int{}, v...)
		res := execSortedSquares(v)
		fmt.Printf("input: %v  -> result: %v\n", input, res)
	}
}

func execSortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	res := make([]int, len(nums))

	pos := len(nums) - 1
	for left <= right {
		res1 := nums[left] * nums[left]
		res2 := nums[right] * nums[right]
		if res1 > res2 {
			res[pos] = res1
			left++
		} else {
			res[pos] = res2
			right--
		}
		pos--
	}

	return res
}
