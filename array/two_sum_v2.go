package array

func TwoSumV2(nums []int, target int) []int {
	res := map[int]int{}

	for index, v := range nums {
		sub := target - v

		val, ok := res[v]
		if !ok {
			res[sub] = index
			continue
		}

		return []int{val, index}
	}	

	return nil
}
