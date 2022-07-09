package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(lenLongestFibSubseq(arr))
}

// lenLongestFibSubseq to get the last num, need two nums
func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	indices := make(map[int]int)
	for i, x := range arr {
		indices[x] = i
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i, x := range arr {
		for j := n - 1; j >= 0 && arr[j]*2 > x; j-- {
			if k, ok := indices[arr[i]-arr[j]]; ok {
				dp[j][i] = max(dp[k][j]+1, 3)
				ans = max(ans, dp[j][i])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
