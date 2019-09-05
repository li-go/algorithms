// https://leetcode.com/problems/reverse-nodes-in-k-group/
package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	head, tail := reverse(head, k)
	for tail != nil {
		h, t := reverse(tail.Next, k)
		tail.Next = h
		tail = t
	}
	return head
}

func reverse(head *ListNode, k int) (*ListNode, *ListNode) {
	nodes := make([]*ListNode, 0, k)
	tmp := head
	for tmp != nil && k > 0 {
		nodes = append(nodes, tmp)
		tmp = tmp.Next
		k--
	}
	if k > 0 {
		return head, nil
	}
	for i := len(nodes) - 1; i > 0; i-- {
		nodes[i].Next = nodes[i-1]
	}
	nodes[0].Next = tmp
	return nodes[len(nodes)-1], nodes[0]
}
