package main

import "fmt"

func main() {
	nums := []int{5, 0, 3, 8, 6}
	fmt.Println(partitionDisjoint(nums))
}

func partitionDisjoint(nums []int) int {
	n := len(nums)
	minRight := make([]int, n)
	minRight[n-1] = nums[n-1]
	for i := n - 2; i > 0; i-- {
		minRight[i] = min(minRight[i+1], nums[i])
	}
	maxLeft := nums[0]
	for i := 1; ; i++ {
		if maxLeft <= minRight[i] {
			return i
		}
		maxLeft = max(maxLeft, nums[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
