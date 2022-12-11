package main

import "fmt"

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
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
