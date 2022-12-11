package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}

// numIslands dfs
// if grid[x][y] == '1', start dfs, during dfs, each '1' will be set to '0'
func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
