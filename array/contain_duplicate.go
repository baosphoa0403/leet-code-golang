package array

import "fmt"

func ContainDuplicate() {
	nums := []int{11, 15, 3, 4, 2, 7, 2}

	rs := execContainDuplicate(nums)
	fmt.Printf("duplicate data = %d\n", rs)
}

func execContainDuplicate(nums []int) int {
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		ele := nums[i]
		if _, ok := seen[ele]; ok {
			return ele
		}

		seen[ele] = ele
	}
	return 0
}
