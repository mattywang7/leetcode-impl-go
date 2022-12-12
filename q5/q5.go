package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println(longestPalindrome2(s))
}

// dp
// when s[i + 1 : j - 1] is palindrome and s[i] == s[j], then s[i : j] is palindrome
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	begin, maxLen := 0, 1
	for L := 2; L <= n; L++ {
		for i := 0; i < n; i++ {
			// j - i + 1 = L
			j := L + i - 1
			if j >= n {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && L > maxLen {
				maxLen = L
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

// we can expand from each edge condition, which means 1 ele or 2 eles
// iterate all center, until we can no longer expand
func longestPalindrome2(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	var expandFromCentre = func(l, r int) int {
		// expand, no way that l > r
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		return (r - 1) - (l + 1) + 1
	}
	start, end := 0, 0
	for i := 0; i < n; i++ {
		len1 := expandFromCentre(i, i)
		len2 := expandFromCentre(i, i+1)
		maxL := max(len1, len2)
		if maxL > (end - start + 1) {
			start = i - (maxL-1)/2
			end = i + maxL/2
		}
	}
	return s[start : end+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
