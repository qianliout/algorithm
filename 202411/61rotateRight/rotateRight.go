package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	node := GenListNode([]int{1, 2, 3, 4, 5})
	PrintListNode(rotateRight(node, 0))
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	cnt := countNode(head)
	k = k % cnt
	if k == 0 {
		return head
	}

	dump := &ListNode{Next: head}
	pre := dump
	cur := dump.Next
	preCnt := 0
	for preCnt < cnt-k {
		pre = pre.Next
		cur = cur.Next
		preCnt++
	}
	pre.Next = nil
	ans := cur
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = dump.Next
	return ans
}

func countNode(head *ListNode) int {
	if head == nil {
		return 0
	}
	return countNode(head.Next) + 1
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return next
}
