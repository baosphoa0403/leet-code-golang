package tree

import "fmt"

func InvertTree() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	printLevelOrder(root)
	out := invertTree(root)
	printLevelOrder(out)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	fmt.Println("before root", root, "left", root.Left, "right", root.Right)

	root.Left, root.Right = root.Right, root.Left
	fmt.Println("after root", root, "left", root.Left, "right", root.Right)

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func printLevelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node.Val, " ")

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}
