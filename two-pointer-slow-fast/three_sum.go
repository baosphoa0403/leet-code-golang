package twopointerslowfast

import (
	"fmt"
	"sort"
)

func ThreeSum() {
	v := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
	}

	for _, v := range v {
		res := threeSum(v)
		fmt.Printf("input: %v  -> result: %v\n", v, res)
	}
}

func threeSum(nums []int) [][]int {
	// important sort
	sort.Ints(nums)
	result := [][]int{}

	// -2 because we need get three element
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
