package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	dump := &ListNode{Next: head}
	cur := head
	cnt := 0
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	p0 := dump
	var pre *ListNode
	cur = p0.Next
	for cnt >= k {
		cnt -= k
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre // 反转
			pre = cur
			cur = nxt
		}
		nxt := p0.Next

		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}
	return dump.Next
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	cnt := 0
	dump := &ListNode{Next: head}
	cur := head
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	var pre *ListNode // 这里一定要是个 nil 节点
	cur = head
	p0 := dump // p0 表示接下来要反转的链表的上一个节点
	for cnt >= k {
		cnt -= k
		for i := 0; i < k; i++ {
			nex := cur.Next
			cur.Next = pre
			pre = cur
			cur = nex
		}

		// 一定要理解这一步
		nex := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nex
	}
	return dump.Next
}
