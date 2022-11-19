package main

import "fmt"

func main() {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain))
}

func largestAltitude(gain []int) int {
	n := len(gain)
	als := make([]int, n+1)
	for i := 0; i < n; i++ {
		als[i+1] = als[i] + gain[i]
	}
	ans := als[0]
	for i := 1; i < n+1; i++ {
		ans = max(ans, als[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
