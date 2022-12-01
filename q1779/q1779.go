package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	points := [][]int{
		{1, 2},
		{3, 1},
		{2, 4},
		{2, 3},
		{4, 4},
	}
	fmt.Println(nearestValidPoint(x, y, points))
}

func nearestValidPoint(x, y int, points [][]int) int {
	ansIdx := -1
	distance := math.MaxInt32
	for i, point := range points {
		if !isValid(point, x, y) {
			continue
		}
		d := abs(x-point[0]) + abs(y-point[1])
		if d < distance {
			ansIdx = i
			distance = d
		}
	}
	return ansIdx
}

func isValid(point []int, x, y int) bool {
	return point[0] == x || point[1] == y
}

// abs write your own simple abs to avoid type transfermation
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
