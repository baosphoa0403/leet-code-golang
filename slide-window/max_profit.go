package slidewindow

import "fmt"

func MaxProfit() {
	inputs := [][]int{
		[]int{7, 1, 5, 3, 6, 4},
	}

	for _, k := range inputs {
		res := maxProfit(k)
		fmt.Printf("input: %v  -> result: %v\n", k, res)
	}
}

func maxProfit(prices []int) int {
	max := 0
	min := 10000

	for _, v := range prices {
		if v < min {
			min = v
		}
		profit := v - min

		if profit > max {
			max = profit
		}
	}

	return max
}
