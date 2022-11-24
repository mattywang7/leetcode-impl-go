package main

import "fmt"

func main() {
	s := "heeellooo"
	words := []string{"hello", "hi", "helo"}
	fmt.Println(expressiveWords(s, words))
}

func expressiveWords(s string, words []string) int {
	ans := 0
	for _, word := range words {
		if canExpand(s, word) {
			ans++
		}
	}
	return ans
}

func canExpand(s, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] != t[j] {
			return false
		}
		ch := s[i]
		cntI := 0
		for i < n && s[i] == ch {
			i++
			cntI++
		}
		cntJ := 0
		for j < m && t[j] == ch {
			j++
			cntJ++
		}
		if cntI < cntJ || (cntI > cntJ && cntI < 3) {
			return false
		}
	}
	return (i == n) && (j == m)
}
