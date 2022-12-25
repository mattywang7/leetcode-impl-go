package main

import "fmt"

var pMap = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]byte, 0)
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pMap[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
