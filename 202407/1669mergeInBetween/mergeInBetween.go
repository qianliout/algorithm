package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	list1 := GenListNode([]int{10, 1, 13, 6, 9, 5})
	list2 := GenListNode([]int{1000000, 1000001, 1000002})
	node := mergeInBetween(list1, 3, 4, list2)
	PrintListNode(node)
}

// 下标从 a 到 b 的全部节点都删除
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dump := &ListNode{Next: list1}
	var nodeA *ListNode
	var nodeB *ListNode
	cur := dump
	idx := 0
	for cur != nil && cur.Next != nil {
		if idx == a {
			nodeA = cur
		}
		if idx-1 == b {
			nodeB = cur.Next
		}
		if nodeA != nil && nodeB != nil {
			break
		}
		idx++
		cur = cur.Next
	}
	nodeA.Next = list2
	last := findEnd(list2)
	last.Next = nodeB
	return dump.Next
}

func findEnd(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return node
	}
	return findEnd(node.Next)
}
