package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	node := GenListNode([]int{1, 2, 3, 3, 4, 4, 5})
	PrintListNode(deleteDuplicates(node))
}

// 迭代的做法
// 什么时候加 dump节点呢，只要头节点也可能删除的时候就需要加 dump 节点
func deleteDuplicates(head *ListNode) *ListNode {
	dump := &ListNode{Next: head}
	cur := dump

	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		v := cur.Next.Val
		if cur.Next.Next.Val != v {
			cur = cur.Next
			continue
		}
		for cur.Next != nil && cur.Next.Val == v {
			cur.Next = cur.Next.Next
		}
	}
	return dump.Next
}

// 只留下不同的数字(如果有相同的数字，则全部删除) 。返回 已排序的链表 。

func deleteDuplicates2(head *ListNode) *ListNode {
	dump := &ListNode{Next: head}
	cur := dump

	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		v := cur.Next.Val
		if cur.Next.Next.Val != v {
			cur = cur.Next
			continue
		}
		for cur.Next.Next != nil && cur.Next.Next.Val == v {
			cur.Next = cur.Next.Next
		}
	}
	return dump.Next
}

// 删除所有重复的元素，使每个元素只出现一次(如果有相同的数据，只保留一个) 。返回 已排序的链表 。
