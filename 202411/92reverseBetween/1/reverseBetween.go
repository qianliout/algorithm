package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	head := GenListNode([]int{1, 2, 3, 4, 5})
	// PrintListNode(reverse(head, nil))
	node := reverseBetween(head, 2, 4)
	PrintListNode(node)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left > 1 {
		head.Next = reverseBetween(head.Next, left-1, right-1)
		return head
	}
	cnt := 0
	rn := head
	for rn != nil && cnt < right {
		rn = rn.Next
		cnt++
	}
	ans := reverse(head, rn)
	return ans
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
