package twopointerslowfast

import (
	"fmt"
	"sort"
)

func FourSum() {

	inputs := map[int][]int{
		0: {1, 0, -1, 0, -2, 2},
		8: {2, 2, 2, 2, 2},
	}

	for k, v := range inputs {
		res := fourSum(v, k)
		fmt.Printf("input: %v  -> result: %v\n", v, res)
	}
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[left] + nums[right] + nums[i] + nums[j]
				if sum == target {
					// fmt.Println("left", left, "right", right, "i", i, "j", j)

					result = append(result, []int{nums[j], nums[i], nums[left], nums[right]})
					left++
					right--

					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}
				if sum < target {
					left++
				}

				if sum > target {
					right--
				}
			}
		}

	}
	return result
}
