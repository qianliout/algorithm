package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nex := head.Next
	for nex != nil && nex.Val == head.Val {
		nex = nex.Next
	}
	netNode := deleteDuplicates(nex)
	head.Next = netNode
	return head
}
