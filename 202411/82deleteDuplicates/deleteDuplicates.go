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
	dump := &ListNode{Next: head}
	cur := head
	for cur.Next != nil && cur.Next.Val == cur.Val {
		cur = cur.Next
	}
	dump.Next = deleteDuplicates(cur.Next)
	return dump.Next
}
