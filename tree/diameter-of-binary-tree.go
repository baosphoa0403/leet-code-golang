package tree

import "fmt"

func DiameterOfBinaryTree() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	out := diameterOfBinaryTree(root)
	fmt.Println(out)
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)

		if leftHeight+rightHeight > maxDiameter {
			maxDiameter = leftHeight + rightHeight
		}

		return 1 + max(leftHeight, rightHeight)
	}

	dfs(root)
	return maxDiameter
}
