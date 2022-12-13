package main

import (
	"fmt"
	"math"
)

func main() {
	x := 123
	fmt.Println(reverse(x))
}

// reverse before push number into stack, check if
// -2^31 <= rev * 10 + digit <= 2^31 - 1
// finally, can be concluded rev must [INT_MIN / 10, INT_MAX / 10]
func reverse(x int) int {
	var ans int
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		ans = ans*10 + digit
		x /= 10
	}
	return ans
}
