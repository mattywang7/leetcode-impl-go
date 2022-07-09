package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	head = reverseKGroup(head, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	hair := &ListNode{Next: head}
	prev := hair
	for head != nil {
		tail := head
		for i := 0; i < k-1; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		// now tail != nil
		tailNext := tail.Next
		head, tail = reverse(head, tail)
		prev.Next = head
		tail.Next = tailNext
		prev = tail
		head = tailNext
	}
	return hair.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	tmp := head
	// how to set the edge condition?
	// when prev.Next == tail, do the last iteration
	for prev != tail {
		next := tmp.Next
		tmp.Next = prev
		// must set prev first,
		// because the node prev is pointing can be abandoned
		prev = tmp
		tmp = next
	}
	return tail, head
}
