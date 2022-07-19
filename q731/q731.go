package main

import "fmt"

func main() {
	m := Constructor()
	fmt.Println(m.Book(10, 20))
	fmt.Println(m.Book(50, 60))
	fmt.Println(m.Book(10, 40))
	fmt.Println(m.Book(5, 15))
	fmt.Println(m.Book(5, 10))
	fmt.Println(m.Book(25, 55))
}

type pair struct {
	start, end int
}

type MyCalendarTwo struct {
	booked, overlaps []pair
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (c *MyCalendarTwo) Book(start, end int) bool {
	for _, p := range c.overlaps {
		if p.start < end && start < p.end {
			return false
		}
	}
	for _, p := range c.booked {
		if p.start < end && start < p.end {
			c.overlaps = append(c.overlaps, pair{max(p.start, start), min(p.end, end)})
		}
	}
	c.booked = append(c.booked, pair{start, end})
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
