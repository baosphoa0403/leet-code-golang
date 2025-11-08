package array

import "fmt"

func ContainsDuplicate() {
	inputs := [][]int{{1, 2, 3, 1}, {1, 2, 3, 4}, {1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}

	for _, v := range inputs {
		res := execContainsDuplicate(v)

		fmt.Printf("input: %v  -> result: %t\n", v, res)
	}
}

func execContainsDuplicate(nums []int) bool {
	set := make(map[int]bool)

	for _, v := range nums {
		_, ok := set[v]
		if !ok {
			set[v] = true
		} else {
			return true
		}
	}

	return false
}
