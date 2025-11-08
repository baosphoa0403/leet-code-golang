package array

import "fmt"

func FindLHS() {
	inputs := [][]int{{1, 3, 2, 2, 5, 2, 3, 7}, {1, 2, 3, 4}, {1, 1, 1, 1}}

	for _, v := range inputs {
		res := execfindLHS(v)
		fmt.Printf("input: %v  -> result: %d\n", v, res)
	}
}

func execfindLHS(nums []int) int {

	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		set[nums[i]] += 1
	}

	maxLength := 0
	for v := range set {
		_, ok := set[v-1]
		if ok {
			maxLength = max(maxLength, set[v]+set[v-1])
		}
	}
	return maxLength
}
