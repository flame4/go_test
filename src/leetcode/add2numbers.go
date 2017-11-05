package leetcode

import (
	_ "fmt"
)
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	follower := head
	followee := follower
	carrier := 0
	for l1 != nil || l2 != nil || carrier > 0 {
		var result int
		var v1 int
		var v2 int
		if l1 != nil {
			v1 = (*l1).Val
			l1 = (*l1).Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = (*l2).Val
			l2 = (*l2).Next
		} else {
			v2 = 0
		}

		result = v1 + v2 + carrier

		(*follower).Val = result % 10
		carrier = result / 10
		(*follower).Next = new(ListNode)
		followee = follower
		follower = (*follower).Next
	}
	(*followee).Next = nil
	return head
}
