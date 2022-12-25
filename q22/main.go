package main

import "fmt"

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	current := make([]byte, 2*n)
	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == len(current) {
			if isValid(current) {
				ans = append(ans, string(current))
			}
		} else {
			current[pos] = '('
			dfs(pos + 1)
			current[pos] = ')'
			dfs(pos + 1)
		}
	}
	dfs(0)
	return ans
}

func isValid(c []byte) bool {
	balance := 0
	for i := range c {
		if c[i] == '(' {
			balance++
		} else {
			balance--
			if balance < 0 {
				return false
			}
		}
	}
	return balance == 0
}
