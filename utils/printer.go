package utils

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PrintTree print the tree by rows
func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			fmt.Printf("%d, ", node.Val)
			queue = queue[1:]
		}
		fmt.Println()
	}
}
