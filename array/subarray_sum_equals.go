package array

func SubarraySum(nums []int, k int) int {
	count := 0
	i := 0
	for i < len(nums) {
		numstmp := nums[i:]
		res := 0

		for _, v := range numstmp {
			res += v
			if res == k {
				count++
			}
		}
		i++
	}

	return count
}

func SubarraySumV2(nums []int, k int) int {
	count := 0
	prefixSum := 0
	freq := map[int]int{}
	freq[0] = 1

	for _, v := range nums {
		prefixSum += v
		need := prefixSum - k

		if times, ok := freq[need]; ok {
			count += times
		}
		freq[prefixSum]++
	}

	return count
}
