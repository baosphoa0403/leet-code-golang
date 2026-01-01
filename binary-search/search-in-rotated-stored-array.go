package binarysearch

import "fmt"

func SearchInRotatedStoredArray() {
	arr := make([]int, 0)
	arr = append(arr, 4, 5, 6, 7, 0, 1, 2)
	res := searchInRotatedStoredArray(arr, 0)
	fmt.Println(res)
}

func searchInRotatedStoredArray(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // left half sorted
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right half sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
