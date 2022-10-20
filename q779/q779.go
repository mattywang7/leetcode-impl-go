package main

import "fmt"

func main() {
	n, k := 2, 2
	fmt.Println(kthGrammar(n, k))
}

func kthGrammar(n, k int) int {
	if n == 1 {
		return 0
	}
	return k&1 ^ 1 ^ kthGrammar(n-1, (k+1)/2)
}
