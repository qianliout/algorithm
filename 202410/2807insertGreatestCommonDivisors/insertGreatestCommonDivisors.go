package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next

	node := &ListNode{Val: bgd(head.Val, next.Val)}
	head.Next = node
	node.Next = insertGreatestCommonDivisors(next)
	return head
}

func lcm(a, b int) int {
	return a * b / bgd(a, b)
}

func bgd(a, b int) int {
	if b == 0 {
		return a
	}
	return bgd(b, a%b)
}
