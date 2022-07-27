package main

import (
	"fmt"
	"unicode"
)

func main() {
	expression := "-1/2+1/2+1/3"
	fmt.Println(fractionAddition(expression))
}

func fractionAddition(expression string) string {
	denominator, numerator := 0, 1
	for i, n := 0, len(expression); i < n; {
		denominator1, sign := 0, 1
		if expression[i] == '-' || expression[i] == '+' {
			if expression[i] == '-' {
				sign = -1
			}
			i++
		}
		for i < n && unicode.IsDigit(rune(expression[i])) {
			denominator1 = denominator1*10 + int(expression[i]-'0')
			i++
		}
		denominator1 = sign * denominator1
		i++

		numerator1 := 0
		for i < n && unicode.IsDigit(rune(expression[i])) {
			numerator1 = numerator1*10 + int(expression[i]-'0')
			i++
		}

		denominator = denominator*numerator1 + denominator1*numerator
		numerator *= numerator1
	}
	if denominator == 0 {
		return "0/1"
	}
	g := gcd(abs(denominator), numerator)
	return fmt.Sprintf("%d/%d", denominator/g, numerator/g)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
