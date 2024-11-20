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
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	// 这里一定得加 fast.Next.Next != nil,不然就会出错，是为啥呢
	// 这里就有点类似二，如果有两个节点，那么应该前面分一个后面分一个
	// 如果前面分两个，后面分0个就会进入死循环，
	// 如果只有 fast.Next!=nil 那么，当只有两个节点时，fast最后就是nil,前面 head 还是两个节点
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	ans := merge(sortList(head), sortList(mid))
	return ans
}

func merge(head1, head2 *ListNode) *ListNode {
	dump := &ListNode{}
	cur := dump
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			cur = cur.Next
			head1 = head1.Next
		} else {
			cur.Next = head2
			cur = cur.Next
			head2 = head2.Next
		}
	}
	if head1 != nil {
		cur.Next = head1
	}
	if head2 != nil {
		cur.Next = head2
	}
	return dump.Next
}
