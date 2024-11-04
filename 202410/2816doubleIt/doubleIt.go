package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func doubleIt(head *ListNode) *ListNode {
	ans := make([]int, 0)
	cur := head
	for cur != nil {
		ans = append(ans, cur.Val)
		cur = cur.Next
	}
	for i, ch := range ans {
		ans[i] = ch * 2
	}
	add := 0
	n := len(ans)

	res := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		v := ans[i] + add
		res[i+1] = v % 10
		add = v / 10
	}
	res[0] = add

	dump := &ListNode{Next: head}
	cur = dump
	i := 0
	for i <= n {
		cur.Val = res[i]
		cur = cur.Next
		i++
	}
	if dump.Val == 0 && dump.Next == nil {
		return nil
	}
	if dump.Val == 0 && dump.Next != nil {
		return dump.Next
	}
	return dump
}
