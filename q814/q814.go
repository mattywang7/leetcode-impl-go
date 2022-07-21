package main

import "matty/leetcode/utils"

func pruneTree(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}

func main() {
	root := &utils.TreeNode{Val: 1}
	root.Right = &utils.TreeNode{Val: 0}
	root.Right.Left = &utils.TreeNode{Val: 0}
	root.Right.Right = &utils.TreeNode{Val: 1}
	root = pruneTree(root)
	utils.PrintTree(root)
}
