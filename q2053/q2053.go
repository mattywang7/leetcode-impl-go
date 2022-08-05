package main

import "fmt"

func main() {
	arr := []string{"d", "b", "c", "b", "c", "a"}
	fmt.Println(kthDistinct(arr, 2))
}

func kthDistinct(arr []string, k int) string {
	m := make(map[string]int)
	for _, s := range arr {
		m[s]++
	}
	for _, s := range arr {
		if m[s] == 1 {
			k--
			if k == 0 {
				return s
			}
		}
	}
	return ""
}
