package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-3, 2, -3, 4, 2}
	fmt.Println(minStartValue(nums))
}

func minStartValue(nums []int) int {
	m := nums[0]
	for _, num := range nums[1:] {
		m = min(m, num)
	}
	return 1 + sort.Search(-m*len(nums), func(startValue int) bool {
		startValue++
		for _, num := range nums {
			startValue += num
			if startValue <= 0 {
				return false
			}
		}
		return true
	})
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
