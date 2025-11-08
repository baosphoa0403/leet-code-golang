package array

import (
	"fmt"
	"sort"
)

func FirstMissingPositive() {
	inputs := [][]int{{1, 2, 0}, {3, 4, -1, 1}, {7, 8, 9, 11, 12}}

	for _, v := range inputs {
		res := execFirstMissingPositive(v)

		fmt.Printf("input: %v  -> result: %d\n", v, res)
	}
}
func execFirstMissingPositive(nums []int) int {
	expected := 1
	sort.Ints(nums)

	for v := range nums {
		if nums[v] == expected {
			expected++
		}
	}

	return expected

}
