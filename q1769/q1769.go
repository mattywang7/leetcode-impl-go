package main

import "fmt"

func main() {
	boxes := "110"
	fmt.Println(minOperations2(boxes))
}

// minOperations1 双重循环
func minOperations1(boxes string) []int {
	ans := make([]int, len(boxes))
	for i := range boxes {
		for j, c := range boxes {
			if c == '1' {
				ans[i] += abs(i - j)
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 当前下标 i 的操作数为 operations_i, i 左侧有 left_i 个球，右侧有 right_i 个球
// 那么下一个下标 i + 1 的操作数 operations_i+1 = operations_i + left_i - right_i
// 因为左侧的球需要多移一步，右侧的球则可以少移一步
func minOperations2(boxes string) []int {
	left := int(boxes[0] - '0')
	right := 0
	operations := 0
	n := len(boxes)
	for i := 1; i < n; i++ {
		if boxes[i] == '1' {
			right++
			operations += i
		}
	}
	ans := make([]int, n)
	ans[0] = operations
	for i := 1; i < n; i++ {
		operations = operations + left - right
		if boxes[i] == '1' {
			left++
			right--
		}
		ans[i] = operations
	}
	return ans
}
