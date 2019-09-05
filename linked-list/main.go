package main

import (
	"fmt"
)

func main() {
	fmt.Println("reverseKGroup:", reverseKGroup(NewListNode([]int{1, 2, 3, 4, 5}), 2).Vals())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Vals() []int {
	head := l
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	return vals
}

func NewListNode(nums []int) *ListNode {
	var head, tail *ListNode
	for _, num := range nums {
		tmp := &ListNode{Val: num}
		if tail == nil {
			head, tail = tmp, tmp
		} else {
			tail.Next, tail = tmp, tmp
		}
	}
	return head
}
