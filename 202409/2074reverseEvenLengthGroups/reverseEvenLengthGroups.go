package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	// head := GenListNode([]int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4})
	// head := GenListNode([]int{1, 1, 0, 6})
	head := GenListNode([]int{5, 2, 6, 3, 9, 1, 7, 3, 8, 46})
	node := reverseEvenLengthGroups(head)
	PrintListNode(node)
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	list := make([][]*ListNode, 0)
	cur := head
	i := 1
	no := make([]*ListNode, 0)
	for cur != nil {
		if len(no) < i {
			no = append(no, cur)
			cur = cur.Next
			// 这一步容易出错
			if cur == nil {
				list = append(list, no)
				break
			}
			continue
		}
		list = append(list, no)
		no = make([]*ListNode, 0)
		i++
	}
	for _, li := range list {
		if len(li)%2 == 0 {
			le, ri := 0, len(li)-1
			for le < ri {
				li[le], li[ri] = li[ri], li[le]
				le++
				ri--
			}
		}
	}
	dump := &ListNode{}
	node := dump
	for _, li := range list {
		for _, n := range li {
			node.Next = n
			node = node.Next
		}
	}
	// 一定得有这一步，不然会有死循环
	node.Next = nil
	return dump.Next
}
