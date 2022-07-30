package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: -8}
	fmt.Println(maxLevelSum(root))
}

func maxLevelSum(root *TreeNode) int {
	ans := 0
	var sum []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if level == len(sum) {
			sum = append(sum, node.Val)
		} else {
			sum[level] += node.Val
		}
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}
	dfs(root, 0)
	for i, s := range sum {
		if s > sum[ans] {
			ans = i
		}
	}
	return ans + 1
}
