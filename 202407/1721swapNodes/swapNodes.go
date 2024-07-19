package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	node := GenListNode([]int{1, 2, 3, 4, 5})
	swapNodes(node, 2)
}

func swapNodes(head *ListNode, k int) *ListNode {
	dump := &ListNode{Next: head}
	var (
		fir *ListNode // 找到的第一个节点的前一个节点
		sec *ListNode
	)
	cur := dump.Next
	for cur != nil {
		if k == 1 {
			fir = cur
			// 找到了 fir，就找另一个了
			cur = dump.Next
			nex := fir
			for nex != nil && nex.Next != nil {
				cur = cur.Next
				nex = nex.Next
			}
			sec = cur
			break
		}
		cur = cur.Next
		k--
	}
	// 只是交换值啊
	if sec == nil || fir == nil {
		return dump.Next
	}
	fir.Val, sec.Val = sec.Val, fir.Val
	return dump.Next
}
