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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dfs func(node1, node2 *ListNode, add int) *ListNode
	dfs = func(node1, node2 *ListNode, add int) *ListNode {
		v1, v2 := 0, 0
		var n1 *ListNode
		var n2 *ListNode
		if node1 == nil && node2 == nil && add == 0 {
			return nil
		}
		if node1 != nil {
			v1 = node1.Val
			n1 = node1.Next
		}
		if node2 != nil {
			v2 = node2.Val
			n2 = node2.Next
		}

		v := v1 + v2 + add
		if v >= 10 {
			v = v % 10
			add = 1
		} else {
			add = 0
		}
		node := &ListNode{Val: v}
		node.Next = dfs(n1, n2, add)
		return node
	}
	return dfs(l1, l2, 0)
}
