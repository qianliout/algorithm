package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	head := GenListNode([]int{1, 2, 3, 4, 5})
	PrintListNode(oddEvenList(head))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd, event := head, head.Next
	ev := event
	for event != nil && event.Next != nil {
		odd.Next = event.Next
		odd = odd.Next
		event.Next = odd.Next
		event = event.Next
	}
	odd.Next = ev
	return head
}
