package main

import "fmt"

func main() {
	intervals := [][]int{
		{1, 3},
		{6, 9},
	}
	newInterval := []int{2, 5}
	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			ans = append(ans, interval)
		} else {
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return ans
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
