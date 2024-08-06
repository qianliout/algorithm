package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func mergeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil && cur.Val == 0 {
		cur = cur.Next
	}
	first := cur
	sum := 0
	for cur != nil && cur.Val != 0 {
		sum += cur.Val
		cur = cur.Next
	}
	if first == nil {
		return nil
	}
	first.Val = sum
	first.Next = mergeNodes(cur)
	return first
}
