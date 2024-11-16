package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	node := GenListNode([]int{1, 2, 3, 3, 4, 4, 5})
	PrintListNode(deleteDuplicates(node))
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
	v := head.Val
	for head != nil && head.Val == v {
		head = head.Next
	}
	return deleteDuplicates(head)
}
