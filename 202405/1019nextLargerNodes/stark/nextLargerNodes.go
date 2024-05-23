package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func nextLargerNodes(head *ListNode) []int {
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	stark := make([]int, 0)
	ans := make([]int, len(list))
	for i := len(list) - 1; i >= 0; i-- {
		for len(stark) > 0 && stark[len(stark)-1] <= list[i] {
			stark = stark[:len(stark)-1]
		}
		if len(stark) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stark[len(stark)-1]
		}
		stark = append(stark, list[i])
	}
	return ans
}
