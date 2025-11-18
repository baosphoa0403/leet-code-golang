package twopointerslowfast

import "fmt"

func RemoveDuplicates() {
	inputs := [][]int{
		[]int{1, 1, 2},
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}

	for _, v := range inputs {
		res := execRemoveDuplicates(v)
		fmt.Printf("input: %v  -> result: %v\n", v, res)
	}

}

func execRemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	fmt.Println("zoo", nums)

	return left + 1
}
