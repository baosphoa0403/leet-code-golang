package array

import "fmt"

func TwoSum() {
	nums := []int{11, 15, 3, 4, 2, 7}
	target := 9
	fmt.Println("TwoSum")
	value1, value2 := execV2(nums, target)
	fmt.Printf("value1 = %d, value2 = %d", value1, value2)
}

func exec(nums []int, target int) (int, int) {
	for i := 0; i < len(nums); i++ {
		ele := nums[i]
		for j := 0; j < len(nums); j++ {
			eleNes := nums[j]
			if ele+eleNes == target {
				return ele, eleNes
			}
		}
		fmt.Println()
	}
	return 0, 0
}

// 2, 7, 11, 15
// 9

// map [7]2
// map
func execV2(nums []int, target int) (int, int) {
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		ele := nums[i]
		if value, ok := seen[ele]; ok {
			return ele, value
		}
		seen[target-ele] = ele
	}

	return -1, -1
}
