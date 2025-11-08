package array

import (
	"fmt"
	"sort"
)

type pair struct {
	num   int
	count int
}

func TopKFrequenceElement() {
	tests := [][]int{
		{1, 2},
		// {1, 1, 1, 2, 2, 3},
		// {1, 1, 1, 2, 2, 3},
	}
	for _, tt := range tests {
		fmt.Printf("%v → %v\n", tt, topKFrequent(tt, 2))
	}
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}

	freqs := make([]pair, 0, len(count))
	for num, c := range count {
		freqs = append(freqs, pair{num, c})
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})

	fmt.Println("freqs", freqs)
	res := make([]int, 0, k)
	for _, item := range freqs {
		res = append(res, item.num)
	}

	return res[:k]
}
