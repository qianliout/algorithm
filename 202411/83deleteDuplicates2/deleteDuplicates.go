package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	node := GenListNode([]int{1, 2, 3, 3, 4, 4, 5})
	PrintListNode(deleteDuplicates(node))
}

func deleteDuplicates1(head *ListNode) *ListNode {
	cur := head

	for cur != nil && cur.Next != nil {
		v := cur.Val
		if cur.Next.Val != v {
			cur = cur.Next
			continue
		}
		for cur.Next != nil && cur.Next.Val == v {
			cur.Next = cur.Next.Next
		}
	}
	return head
}

// 删除所有重复的元素，使每个元素只出现一次(如果有相同的数据，只保留一个) 。返回 已排序的链表 。

// 一定要和82题比较

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Val != head.Val {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
	v := head.Val
	nex := head.Next
	for nex != nil && nex.Val == v {
		nex = nex.Next
	}
	head.Next = deleteDuplicates(nex)
	return head
}
