package main

import (
	"math"

	. "outback/algorithm/common/listnode"
)

func main() {

}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	inf := math.MaxInt32
	a, b, pre := inf, -1, -1
	var preNode *ListNode
	mi := inf
	cur := head
	start := 0
	for cur != nil && cur.Next != nil {
		if preNode != nil && cur.Next != nil {
			if (preNode.Val < cur.Val && cur.Val > cur.Next.Val) || (preNode.Val > cur.Val && cur.Val < cur.Next.Val) {
				a = min(a, start)
				b = max(b, start)
				if pre != -1 {
					mi = min(mi, start-pre)
				}
				pre = start
			}
		}
		start++
		preNode = cur
		cur = cur.Next
	}
	if mi == inf || a == inf || b == -1 || a == b {
		return []int{-1, -1}
	}
	return []int{mi, b - a}
}
