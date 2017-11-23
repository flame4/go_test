package leetcode

import (
	"fmt"
	"testing"
)

func TestAdd2Numbers_1(t *testing.T) {
	l1 := makeList([]int{1, 2, 3, 4})
	l2 := makeList([]int{1, 2, 3, 4})
	l3 := makeList([]int{2, 4, 6, 8})
	validateList(addTwoNumbers(l1, l2), l3, t)
}

func TestUse(t *testing.T) {
	l1 := makeList([]int{1, 2, 3, 4})
	ShowList(l1)
}

func makeList(l []int) *ListNode {
	head := new(ListNode)
	follower := head
	for index, val := range l {
		(*follower).Val = val
		if index < len(l)-1 {
			(*follower).Next = new(ListNode)
		}
		follower = (*follower).Next
	}
	return head
}

func validateList(l1 *ListNode, l2 *ListNode, t *testing.T) {
	//ShowList(l1)
	//ShowList(l2)
	for (l1 != nil) || (l2 != nil) {
		//fmt.Println(1)
		switch {
		case l1 == nil && l2 != nil:
			t.Fatal("List Length Not Equal! l2 larger!")
		case l2 == nil && l1 != nil:
			t.Fatal("List Length Not Equal! l1 larger")
		}
		if (*l1).Val != (*l2).Val {
			t.Logf("l1=%d, l2=%d \n", (*l1).Val, (*l2).Val)
			t.Fatal("l1.val != l2.val!")
		}
		l1, l2 = (*l1).Next, (*l2).Next
	}
}

func ShowList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d,", (*l).Val)
		l = (*l).Next
	}
	fmt.Println()
}
