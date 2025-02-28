package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left > 1 {
		head.Next = reverseBetween(head.Next, left-1, right-1)
		return head
	}
	// find right.Next Node
	end := head
	// 题目保证了 right 在链表长度内
	for right > 0 && end != nil {
		end = end.Next
		right--
	}
	return reverse(head, end)
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
