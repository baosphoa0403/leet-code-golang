package twopointerslowfast

import "fmt"

func RemoveElement() {
	inputs := map[int][]int{
		3: []int{3, 2, 2, 3},
		2: []int{0, 1, 2, 2, 3, 0, 4, 2},
	}

	for k, v := range inputs {
		res := removeElement(v, k)
		fmt.Printf("input: %v  -> result: %v\n", v, res)
	}

}

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
