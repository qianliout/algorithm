package main

import (
	. "outback/algorithm/common/listnode"
	"sort"
)

func main() {
	head := GenListNode([]int{1, 2, 3, 4})
	reorderList(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return next
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	cnt := 0
	a, b := headA, headB
	for cnt < 3 {
		// 一定要先判断，因为可能最开始的点就是相交点
		if a == b {
			return a
		}
		a = a.Next
		if a == nil {
			a = headB
			cnt++
		}
		b = b.Next
		if b == nil {
			b = headA
			cnt++
		}

	}
	return nil
}

func reorderList(head *ListNode) {
	var dfs func(root *ListNode) *ListNode
	dfs = func(root *ListNode) *ListNode {
		if root == nil || root.Next == nil || root.Next.Next == nil {
			return root
		}
		fir, sec := root, root.Next
		pre, end := root, root.Next
		for end != nil && end.Next != nil {
			pre = pre.Next
			end = end.Next
		}

		pre.Next = nil
		nex := dfs(sec)
		root.Next = end
		end.Next = nex
		return fir
	}
	dfs(head)
}

func sortList(head *ListNode) *ListNode {
	data := make([]*ListNode, 0)
	for head != nil {
		data = append(data, head)
		head = head.Next
	}
	sort.Slice(data, func(i, j int) bool { return data[i].Val < data[j].Val })
	dump := &ListNode{}
	cur := dump
	for i := 0; i < len(data); i++ {
		cur.Next = data[i]
		cur = cur.Next
	}
	// 这一步最重要，不然会死循环
	cur.Next = nil
	return dump.Next
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 这样做的目的是怕第一个节点就是环的入口的
	dumy := &ListNode{Next: head}
	slow, fast := dumy, dumy
	for {
		if slow == nil || fast == nil || fast.Next == nil {
			return nil // 说明无环
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = dumy
			break
		}
	}
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func isPalindrome(head *ListNode) bool {
	data := make([]int, 0)
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	le, ri := 0, len(data)-1
	for le < ri {
		if data[le] != data[ri] {
			return false
		}
		le++
		ri--
	}
	return true
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	return merge2List(left, right)
}

func merge2List(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val < b.Val {
		a.Next = merge2List(a.Next, b)
		return a
	}
	b.Next = merge2List(a, b.Next)
	return b
}
