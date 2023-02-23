package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 5, 0, 3, 5}
	fmt.Println(minimumOperations(nums))
}

// greedy, choose the min positive number in the array
func minimumOperations(nums []int) int {
	sort.Ints(nums)
	cnt := 0
	for nums[len(nums)-1] != 0 {
		var toMinus int
		for _, num := range nums {
			if num != 0 {
				toMinus = num
				break
			}
		}
		if toMinus == 0 {
			return 0
		}
		for i := range nums {
			if nums[i] > 0 {
				nums[i] -= toMinus
			}
		}
		cnt++
	}
	return cnt
}
