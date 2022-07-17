package main

import "fmt"

func main() {
	nums := []int{5, 4, 0, 3, 1, 6, 2}
	fmt.Println(arrayNesting(nums))
}

func arrayNesting(nums []int) int {
	ans := 0
	vis := make([]bool, len(nums))
	for i := range vis {
		cnt := 0
		for !vis[i] {
			vis[i] = true
			i = nums[i]
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}
