package tree

import "fmt"

func PreorderTraversal() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	out := preOrder(root)
	fmt.Println(out)
}

func preOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	res = append(res, root.Val)
	res = append(res, preOrder(root.Left)...)
	res = append(res, preOrder(root.Right)...)

	return res
}
