package tree

import "fmt"

func InorderTraversal() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	out := inOrder(root)
	fmt.Println(out)
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	res = append(res, inOrder(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inOrder(root.Right)...)
	return res
}
