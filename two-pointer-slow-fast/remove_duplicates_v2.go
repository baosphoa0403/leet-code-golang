package twopointerslowfast

// nums := []int{0,0,1,1,1,2,2,3,3,4}
func RemoveElementV2(nums []int, val int) int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
