package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(5)
	obj.Push(7)
	obj.Push(5)
	obj.Push(7)
	obj.Push(4)
	obj.Push(5)

	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
}

type FreqStack struct {
	freq    map[int]int
	group   map[int][]int
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:    make(map[int]int),
		group:   make(map[int][]int),
		maxFreq: 0,
	}
}

func (f *FreqStack) Push(Val int) {
	f.freq[Val]++
	f.group[f.freq[Val]] = append(f.group[f.freq[Val]], Val)
	f.maxFreq = max(f.maxFreq, f.freq[Val])
}

func (f *FreqStack) Pop() int {
	lenGroupMaxFreq := len(f.group[f.maxFreq])
	val := f.group[f.maxFreq][lenGroupMaxFreq-1]
	f.freq[val]--
	f.group[f.maxFreq] = f.group[f.maxFreq][:lenGroupMaxFreq-1]
	if len(f.group[f.maxFreq]) == 0 {
		f.maxFreq--
	}
	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
