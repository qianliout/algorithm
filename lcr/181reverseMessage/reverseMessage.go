package main

import (
	. "outback/algorithm/common/listnode"
	"strings"
)

func main() {

}

func reverseMessage(message string) string {
	ss := strings.Split(message, " ")
	le, ri := 0, len(ss)-1
	for le < ri {
		ss[le], ss[ri] = ss[ri], ss[le]
		le++
		ri--
	}
	ss2 := make([]string, 0)
	for _, ch := range ss {
		if ch != "" {
			ss2 = append(ss2, ch)
		}
	}

	ans := strings.Join(ss2, " ")
	return ans
}

func dynamicPassword(password string, target int) string {
	ss := []byte(password)
	pre := ss[:target]
	ans := append(ss[target:], pre...)
	return string(ans)
}

func deleteNode2(head *ListNode, val int) *ListNode {
	dump := &ListNode{Next: head}
	cur := dump
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return dump.Next
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode(head.Next, val)
	return head
}
