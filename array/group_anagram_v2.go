package array

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{strs}
	}

	res := map[string][]string{}

	for _, v := range strs {
		key := sortText(v)
		_, ok := res[key]
		if !ok {
			res[key] = make([]string, 0)
			res[key] = append(res[key], v)
			continue
		}
		res[key] = append(res[key], v)
	}

	res1 := [][]string{}
	for _, v := range res {
		res1 = append(res1, v)
	}

	return res1
}

func sortText(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
