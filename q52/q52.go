package main

import "fmt"

func main() {
	fmt.Println(totalNQueens(10))
}

func totalNQueens(n int) int {
	ans := 0
	columns := make([]bool, n)
	diagonals1 := make([]bool, 2*n-1)
	diagonals2 := make([]bool, 2*n-1)
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			ans++
			return
		}
		for col, hasQueen := range columns {
			d1, d2 := row+n-1-col, row+col
			if hasQueen || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			columns[col] = true
			diagonals1[d1] = true
			diagonals2[d2] = true
			backtrack(row + 1)
			columns[col] = false
			diagonals1[d1] = false
			diagonals2[d2] = false
		}
	}
	backtrack(0)
	return ans
}
