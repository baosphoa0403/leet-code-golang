package tree

import "fmt"

func IsSubtree() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	// root.Left.Right.Left = &TreeNode{Val: 0}

	subRoot := &TreeNode{Val: 4}
	subRoot.Left = &TreeNode{Val: 1}
	subRoot.Right = &TreeNode{Val: 2}

	out := isSubtree(root, subRoot)
	fmt.Println(out)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
