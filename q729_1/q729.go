package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"math"
)

// 对于给定的区间 [start, end), 每次查找起点 >= end 的第一个区间 [l1, r1),
// 同时紧挨着 [l1, r1) 的前一个区间 [l2, r2),
// 如果满足 r1 <= start < end <= l1, 则该区间可以预定。

type MyCalender struct {
	*redblacktree.Tree // 保持元素排序和支持快速插入
}

func Constructor() MyCalender {
	t := redblacktree.NewWithIntComparator()
	t.Put(math.MaxInt32, nil) // pivot, simplify code
	return MyCalender{t}
}

func (c MyCalender) Book(start, end int) bool {
	node, _ := c.Ceiling(end) // first l1 >= end
	it := c.IteratorAt(node)
	if !it.Prev() || it.Value().(int) <= start {
		c.Put(start, end)
		return true
	}
	return false
}

func main() {
	c := Constructor()
	// equal to c.Book directly, the & operation is performed automatically
	fmt.Println(c.Book(10, 20))
	fmt.Println(c.Book(15, 25))
}
