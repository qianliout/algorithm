package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	node := GenListNode([]int{1, 2, 3, 4, 5})
	PrintListNode(reverseKGroup(node, 3))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	cnt := 0
	dump := &ListNode{Next: head}
	pre := dump
	cur := head
	for cur != nil && cnt < k {
		pre = pre.Next
		cur = cur.Next
		cnt++
	}
	if cnt < k {
		return head
	}

	pre.Next = nil
	next := reverseKGroup(cur, k)

	ans := reverse(head, nil)
	cur1 := ans
	for cur1 != nil && cur1.Next != nil {
		cur1 = cur1.Next
	}
	cur1.Next = next
	return ans
}

func cntNode(head *ListNode) int {
	cnt := 0
	cur := head
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	return cnt
}
func reverse(head *ListNode, end *ListNode) *ListNode {
	if head == nil || head.Next == nil || head == end || head.Next == end {
		return head
	}
	next := reverse(head.Next, end)
	head.Next.Next = head
	head.Next = end
	return next
}
