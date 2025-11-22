package slidewindow

import "fmt"

func FindMaxAverage() {
	inputs := map[int][]int{
		4: {1, 12, -5, -6, 50, 3},
	}

	for k, v := range inputs {
		res := findMaxAverage(v, k)
		fmt.Printf("input: %v  -> result: %v\n", v, res)
	}
}

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	left := 0
	for right := k; right < len(nums); right++ {
		fmt.Println("sum", sum, "nums[right]", nums[right], "nums[left]", nums[left])
		sum += nums[right]
		sum -= nums[left]
		left++

		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
