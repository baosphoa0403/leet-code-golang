package array

import (
	"fmt"
	"strings"
)

func IsSubsequence() {
	inputs := [][]string{{"b", "abc"}}

	for _, item := range inputs {
		res := execIsSubsequence(item[0], item[1])
		fmt.Printf("input1: %s - input2: %s -> result: %t\n", item[0], item[1], res)
	}
}

func execIsSubsequence(s string, t string) bool {
	arrTmp := strings.Split(s, "")
	if len(arrTmp) == 0 {
		return true
	}

	arrTmp2 := strings.Split(t, "")
	j := 0

	for i := 0; i < len(arrTmp2); i++ {
		if j == len(arrTmp) {
			break
		}

		if arrTmp2[i] == arrTmp[j] {
			j++
		}
	}

	return j == len(arrTmp)
}
