package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(twoEggDrop(2))
}

// dp[i][j] - when you have (i + 1) eggs, the min steps to confirm jth floor
// refer: https://leetcode.cn/problems/egg-drop-with-2-eggs-and-n-floors/solution/dong-tai-gui-hua-shu-xue-tui-dao-by-tang-1zz1/
func twoEggDrop(n int) int {
	var dp [][]int
	tmp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		tmp[i] = i
	}
	dp = append(dp, tmp)
	dp = append(dp, make([]int, len(tmp)))
	for i := 0; i < len(tmp); i++ {
		dp[1][i] = math.MaxInt32
	}
	dp[1][0] = 0
	for j := 1; j <= n; j++ {
		for k := 1; k <= j; k++ {
			dp[1][j] = min(dp[1][j], max(dp[0][k-1]+1, dp[1][j-k]+1))
		}
	}
	return dp[1][n]
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
