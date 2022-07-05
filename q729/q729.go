package main

import "fmt"

type pair struct {
	start, end int
}

type MyCalendar []pair

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start, end int) bool {
	for _, p := range *c {
		if p.start < end && p.end > start {
			return false
		}
	}
	*c = append(*c, pair{start: start, end: end})
	return true
}

func main() {
	c := MyCalendar{}
	// equal to c.Book directly, the & operation is performed automatically
	fmt.Println((&c).Book(10, 20))
	fmt.Println(c.Book(15, 25))
}
