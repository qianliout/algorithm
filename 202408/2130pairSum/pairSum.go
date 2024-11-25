package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func pairSum(head *ListNode) int {
	cnt := make([]int, 0)
	cur := head
	for cur != nil {
		cnt = append(cnt, cur.Val)
		cur = cur.Next
	}
	// 1 <= Head.val <= 105
	mx := 0
	n := len(cnt)
	for i := 0; i < n/2; i++ {
		mx = max(mx, cnt[i]+cnt[n-i-1])
	}
	return mx
}
