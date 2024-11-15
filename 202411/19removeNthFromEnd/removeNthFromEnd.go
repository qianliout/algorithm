package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	node := GenListNode([]int{1, 2, 3, 4, 5})
	PrintListNode(removeNthFromEnd(node, 20))
	// PrintListNode(removeNthFromEnd(node, 1))
	// PrintListNode(removeNthFromEnd(node, 5))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cnt := cntNode(head)

	var dfs func(pre *ListNode, n int)
	dfs = func(pre *ListNode, n int) {
		if pre == nil {
			return
		}
		if n == 0 {
			if pre != nil && pre.Next != nil {
				pre.Next = pre.Next.Next
			}
			return
		}
		dfs(pre.Next, n-1)
	}
	dump := &ListNode{Next: head}
	dfs(dump, cnt-n)
	return dump.Next
}

func cntNode(head *ListNode) int {
	if head == nil {
		return 0
	}
	return cntNode(head.Next) + 1
}
