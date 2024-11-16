package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	node := GenListNode([]int{1, 4, 3, 2, 5, 2})
	PrintListNode(partition(node, 3))
}

func partition(head *ListNode, x int) *ListNode {
	mx, mi := &ListNode{}, &ListNode{}
	le, ri := mi, mx
	for head != nil {
		if head.Val < x {
			le.Next = head
			le = le.Next
		} else {
			ri.Next = head
			ri = ri.Next
		}
		head = head.Next
	}

	ri.Next = nil // 这一步是重点，为啥呢,因为直接使用的 head,head 后面还会有Next，接到其他节点上了
	le.Next = mx.Next

	return mi.Next
}
