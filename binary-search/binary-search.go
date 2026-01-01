package binarysearch

import "fmt"

func BinarySearch() {
	arr := make([]int, 0)
	arr = append(arr, -1, 0, 3, 5, 9, 12)
	res := search(arr, 2)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
