package tree

import "fmt"

func PostorderTraversal() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	out := postOrder(root)
	fmt.Println(out)
}

func postOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	res = append(res, postOrder(root.Left)...)
	res = append(res, postOrder(root.Right)...)
	res = append(res, root.Val)
	return res
}
