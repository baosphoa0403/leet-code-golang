package twopointerslowfast

import "fmt"

func RemoveElement() {
	inputs := map[int][]int{
		3: []int{3, 2, 2, 3},
		2: []int{0, 1, 2, 2, 3, 0, 4, 2},
	}

	for k, v := range inputs {
		res := execRemoveElement(v, k)
		fmt.Printf("input: %v  -> result: %v\n", v, res)
	}

}

func execRemoveElement(nums []int, val int) int {
	left := 0
	// for left < c-1 {
	// 	if nums[left] == val {
	// 		fmt.Println("nums[left]", nums[left], nums[left+1])
	// 		nums[left] = nums[left+1]
	// 	}
	// 	left++
	// }
	for i := 1; i < len(nums); i++ {
		if nums[left] != val {
			nums[left] = nums[i]
			fmt.Println("nums", nums)
			left++
		}
	}

	// fmt.Println("nums", nums)
	return len(nums) - left
}
