package twopointerslowfast

import "fmt"

func MoveZeroes() {
	// num1s := []int{0, 1, 0, 3, 12}
	num1s := []int{1, 0, 2, 3}

	slow := 0
	fast := 0

	for fast < len(num1s) {
		if num1s[fast] != 0 {
			num1s[slow] = num1s[fast]

			if slow != fast {
				num1s[fast] = 0
			}

			slow++
		}
		fast++
	}

	fmt.Println("num1s: ", num1s)

}
