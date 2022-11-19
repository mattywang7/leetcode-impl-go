package main

import (
	"fmt"
	"math"
)

func main() {
	poured, queryRow, queryGlass := 2, 1, 1
	fmt.Println(champagneTower(poured, queryRow, queryGlass))
}

// i and j both start from 0
func champagneTower(poured int, query_row int, query_glass int) float64 {
	row := []float64{float64(poured)}
	for i := 1; i <= query_row; i++ {
		nextRow := make([]float64, i+1)
		for j, volume := range row {
			if volume > 1 {
				nextRow[j] += (volume - 1) / 2
				nextRow[j+1] += (volume - 1) / 2
			}
		}
		row = nextRow
	}
	return math.Min(1.0, row[query_glass])
}
