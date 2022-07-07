package main

import "fmt"

func main() {
	position := []int{2, 2, 2, 3, 3}
	fmt.Println(minCostToMoveChips(position))
}

func minCostToMoveChips(position []int) int {
	cnt := [2]int{}
	for _, p := range position {
		cnt[p%2]++
	}
	return min(cnt[0], cnt[1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
