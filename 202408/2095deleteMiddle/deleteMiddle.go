package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	list := GenListNode([]int{1, 3, 4, 7, 1, 2, 6})
	deleteMiddle(list)
	PrintListNode(list)
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	n := count(head)
	dump := &ListNode{Next: head}
	pre := dump
	cur := head
	i := 0
	for i < n/2 && cur != nil {
		pre = cur
		cur = cur.Next
		i++
	}
	if cur != nil {
		pre.Next = cur.Next
	}

	return dump.Next
}

func count(head *ListNode) int {
	if head == nil {
		return 0
	}
	a := 1 + count(head.Next)
	return a
}
