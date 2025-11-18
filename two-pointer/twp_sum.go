package twopointer

import "fmt"

func TwoSum() {
	inputs := map[int][]int{
		9:   []int{2, 7, 11, 15},
		6:   []int{2, 3, 4},
		-1:  []int{-1, 0},
		100: []int{5, 25, 75},
		8:   []int{1, 2, 3, 4, 4, 9, 56, 90},
	}

	for k, v := range inputs {
		res := ExectwoSum(v, k)
		fmt.Printf("input: %v  -> result: %v\n", v, res)
	}

}

func ExectwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum > target {
			right--
		}

		if sum < target {
			left++
		}
	}

	return nil
}
