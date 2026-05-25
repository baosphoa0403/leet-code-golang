package array

import (
	"sort"
)

func TopKFrequentV2(nums []int, k int) []int {
	res := map[int]int{}
	for _, v := range nums {
		res[v]++
	}

	type pair struct {
		number int
		count  int
	}

	pairs := []pair{}
	for k, v := range res {
		pairs = append(pairs, pair{
			number: k,
			count:  v,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	res1 := []int{}
	for _, v := range pairs {
		res1 = append(res1, v.number)
	}

	return res1[:k]
}
