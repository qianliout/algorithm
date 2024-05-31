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
func numComponents1(head *ListNode, nums []int) int {
	exit := map[int]bool{}
	for _, ch := range nums {
		exit[ch] = true
	}
	ans := 0
	for head != nil {
		if exit[head.Val] {
			ans++
			for head != nil && exit[head.Val] {
				head = head.Next
			}
		} else {
			head = head.Next
		}
	}
	return ans
}

func numComponents(head *ListNode, nums []int) int {
	exit := map[int]bool{}
	for _, ch := range nums {
		exit[ch] = true
	}
	ans := 0
	for head != nil {
		if !exit[head.Val] {
			head = head.Next
			continue
		}
		ans++
		for head != nil && exit[head.Val] {
			head = head.Next
		}
	}
	return ans
}
