package array

func MaxSubArray(nums []int) int {
	localMax := nums[0]
	globalMax := nums[0]

	for i := 1; i < len(nums); i++ {
		v := nums[i]

		localMax = max(localMax+v, v)
		globalMax = max(globalMax, localMax)
	}
	return globalMax
}
