package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dump := &ListNode{Next: head}
	p0 := dump
	var pre *ListNode
	// 题目中的数据保证了 left,right 在链表范围内，所以可以不用判断
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}
	cur := p0.Next
	// 这里循环次数是 i<=right - left,也是容易出错的点
	// 怎么理解呢：当 left==right==1的时候,也是需要反转的，此时的目的是为了更新 cur和 pre
	// 不然后面更新 p0的时候，会导致链表断开
	// 可以使用小数据量进行测试一下
	for i := 0; i <= right-left; i++ {
		nex := cur.Next
		cur.Next = pre // 反转
		pre = cur
		cur = nex
	}
	p0.Next.Next = cur
	p0.Next = pre
	return dump.Next
}
