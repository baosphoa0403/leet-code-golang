package tree

import (
	"fmt"
	"math"
)

func IsBalanced() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	out := isBalanced(root)
	fmt.Println(out)
}

func isBalanced(root *TreeNode) bool {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if left == -1 || right == -1 || math.Abs(float64(left-right)) > 1 {
			return -1
		}

		return 1 + max(left, right)
	}

	return dfs(root) != -1
}
