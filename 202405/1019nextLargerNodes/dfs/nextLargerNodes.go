package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func nextLargerNodes(head *ListNode) []int {
	mem := make(map[*ListNode][]int)
	return dfs(head, mem)
}

func dfs(head *ListNode, mem map[*ListNode][]int) []int {
	ans := make([]int, 0)
	if head == nil {
		return ans
	}
	if val, ok := mem[head]; ok {
		return val
	}
	if head.Next == nil {
		ans = append(ans, 0)
		mem[head] = ans
		return ans
	}

	nex := head.Next
	for nex != nil {
		if nex.Val > head.Val {
			ans = append(ans, nex.Val)
			break
		}
		nex = nex.Next
	}
	if len(ans) == 0 {
		ans = append(ans, 0)
	}
	next := dfs(head.Next, mem)
	ans = append(ans, next...)

	mem[head] = ans
	return ans
}
