package main

import "fmt"

func main() {
	s := "0100"
	fmt.Println(minOperations(s))
}

func minOperations(s string) int {
	cnt := 0
	for i, c := range s {
		if i%2 != int(c-'0') {
			cnt++
		}
	}
	return min(cnt, len(s)-cnt)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
