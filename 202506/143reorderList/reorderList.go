package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {
	head := GenListNode([]int{1, 2, 3, 4, 5})
	reorderList(head)
	PrintListNode(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList1(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	fir, sec := head, head.Next
	last := head
	for last != nil && last.Next != nil && last.Next.Next != nil {
		last = last.Next
	}
	if last != nil && last.Next != nil {
		fir.Next = last.Next
		fir.Next.Next = sec
		last.Next = nil
	}
	reorderList1(sec)
}

func reorderList2(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	dump := head
	end1 := dump
	end2 := dump.Next

	for end2 != nil && end2.Next != nil {
		end1 = end1.Next
		end2 = end2.Next
	}
	end1.Next = nil
	first := head.Next
	reorderList2(first)

	head.Next = end2
	head.Next.Next = first
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	fir, sec := head, head.Next
	last := head
	end := head.Next
	for end != nil && end.Next != nil {
		last = last.Next
		end = end.Next
	}
	fir.Next = last.Next
	fir.Next.Next = sec
	last.Next = nil
	reorderList(sec)
}
