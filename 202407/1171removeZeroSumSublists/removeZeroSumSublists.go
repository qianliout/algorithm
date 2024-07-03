package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: -1}
	removeZeroSumSublists(head)
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	dump := &ListNode{Next: head}
	sum := 0
	cnt := make(map[int]*ListNode)
	cnt[0] = dump
	for cur := dump.Next; cur != nil; cur = cur.Next {
		sum += cur.Val
		cnt[sum] = cur
	}
	cur := dump
	sum = 0
	for cur != nil {
		sum += cur.Val
		cur.Next = cnt[sum].Next
		cur = cur.Next
	}
	return dump.Next
}
