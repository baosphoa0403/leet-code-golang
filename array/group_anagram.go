package array

import "fmt"

func GroupAnagram() {
	tests := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
	}

	for _, v := range tests {
		fmt.Printf("value: %s -> res: %#v", v, isAnagramGroup(v))
	}
}

func isAnagramGroup(words []string) map[string][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		count := make([]int, 26)

		for _, ch := range word {
			count[ch-'a']++
		}

		key := fmt.Sprint(count)
		groups[key] = append(groups[key], word)
	}

	return groups
}
