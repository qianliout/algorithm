package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	d := k
	last := head
	for last != nil && d > 0 {
		last = last.Next
		d--
	}
	if d > 0 {
		return head
	}

	node := reverse(head, last)

	nex := reverseKGroup(last, k)

	cur := node
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = nex
	return node
}

func reverse(head, end *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next == end {
		return head
	}
	nex := reverse(head.Next, end)
	head.Next.Next = head
	head.Next = nil
	return nex
}
