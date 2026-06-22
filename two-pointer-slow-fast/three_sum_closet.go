package twopointerslowfast

import (
	"sort"
)

// nums := []int{-1, 2, 1, -4}
// target := 1
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-1; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}
			if sum < target {
				left++
			}

			if sum > target {
				right--
			}
			if sum == target {
				return sum
			}
		}
	}

	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
