package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for i, v := range nums {
		remain := target - v
		if idx, ok := table[remain]; ok {
			return []int{idx, i}
		}
		table[v] = i
	}
	return nil
}
