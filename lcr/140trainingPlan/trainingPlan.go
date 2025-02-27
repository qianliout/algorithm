package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func trainingPlan1(head *ListNode, cnt int) *ListNode {
	// 方法一，反转链表，然后找正数
	// 找到之后再把原链表再反转，然后放回
	// 容易出错的点是再次反转，因为不再次反转，那么链表的结构会变动
	head = reverse(head)
	cur := head
	cnt--
	for cnt > 0 && cur != nil {
		cur = cur.Next
		cnt--
	}
	reverse(head)

	return cur
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nxt := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return nxt
}

// 直接计算，这里的边界条件不好判断，此时可以建设 cnt=1 或head 只有一个节点
func trainingPlan(head *ListNode, cnt int) *ListNode {
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	sub := n - cnt
	if sub < 0 {
		return nil
	}
	cur = head
	for sub > 0 && cur != nil {
		sub--
		cur = cur.Next
	}
	return cur
}
