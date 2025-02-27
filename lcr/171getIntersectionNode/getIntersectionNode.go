package main

import (
	"fmt"
	. "outback/algorithm/common/listnode"
)

func main() {
	fmt.Println("vim-go")
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	cnt := 0
	for cnt <= 3 {
		if a == b {
			return a
		}

		a = a.Next
		if a == nil {
			a = headB
			cnt++
		}
		b = b.Next
		if b == nil {
			b = headA
			cnt++
		}
	}
	return nil
}

// 这样做是错的，是为啥呢
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	cnt := 0
	// 测试后和这里的 cnt 无关
	for cnt <= 7 {
		if a == nil {
			a = headB
			cnt++
		}
		if b == nil {
			b = headA
			cnt++
		}
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}
