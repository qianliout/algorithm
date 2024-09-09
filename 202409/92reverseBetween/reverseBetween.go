package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	head := GenListNode([]int{1, 2, 3, 4, 5})
	PrintListNode(reverseBetween(head, 2, 4))
}

// 请你反转从位置 left 到位置 right 的链表节点
// func reverseBetween2(head *ListNode, left int, right int) *ListNode {
// 	dummy := &ListNode{Next: head}
// 	cur := dummy
// 	var reverse func(head *ListNode) *ListNode
// 	var rightN *ListNode
// 	reverse = func(head *ListNode) *ListNode {
// 		if head == nil || head.Next == nil || head.Val == right {
// 			if head.Val == right {
// 				rightN = head.Next
// 			}
// 			return head
// 		}
// 		next := reverse(head.Next)
//
// 		head.Next.Next = head
// 		head.Next = nil
// 		return next
// 	}
//
// 	for cur != nil && cur.Next != nil {
// 		if cur.Next.Val == left {
// 			next := reverse(cur.Next)
// 			cur.Next = next
// 			break
// 		}
// 		cur = cur.Next
// 	}
// 	cur = dummy
// 	for cur != nil && cur.Next != nil {
// 		cur = cur.Next
// 	}
// 	cur.Next = rightN
// 	return dummy.Next
// }

// 请你反转从位置 left 到位置 right 的链表节点
// func reverseBetween(head *ListNode, left int, right int) *ListNode {
// 	if head == nil || head.Next == nil || right <= left || left < 1 {
// 		return head
// 	}
//
// 	if left == 1 {
// 		rightNode := head
// 		for rightNode != nil && right > 0 {
// 			rightNode = rightNode.Next
// 			right--
// 		}
// 		lst := reverse(head, rightNode)
//
// 		cur := lst
// 		for cur != nil && cur.Next != nil {
// 			cur = cur.Next
// 		}
// 		cur.Next = rightNode
// 		return lst
// 	}
// 	head.Next = reverseBetween(head.Next, left-1, right-1)
// 	return head
// }
//
// func reverse(head, end *ListNode) *ListNode {
// 	if head == nil || head.Next == nil || head.Next == end {
// 		return head
// 	}
// 	nex := reverse(head.Next, end)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return nex
// }

// 反转从位置 left 到位置 right 的链表节点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || right <= left || left < 1 {
		return head
	}
	if left == 1 {
		rn := head
		for rn != nil && right > 0 {
			rn = rn.Next
			right--
		}
		next := reverse(head, rn)
		cur := next
		for cur != nil && cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = rn
		return next
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func reverse(head, end *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next == end {
		return head
	}
	next := reverse(head.Next, end)
	head.Next.Next = head
	head.Next = nil
	return next
}
