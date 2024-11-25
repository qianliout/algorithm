package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	head := GenListNode([]int{1, 2, 3, 4, 5})
	reverseList3(head)
}

// 递归解法
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return next
}

// 迭代
func reverseList(head *ListNode) *ListNode {
	// dump := &ListNode{Next: head}
	var pre *ListNode
	cur := head
	for cur != nil {
		nex := cur.Next
		cur.Next = pre // 反转
		pre = cur
		cur = nex
	}
	return pre
}

func reverseList3(head *ListNode) *ListNode {
	dump := &ListNode{Next: head}
	// 这样写会进入死循环，pre.Next 必须是 nil
	pre := dump
	cur := head
	for cur != nil {
		nex := cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}
