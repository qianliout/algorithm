package main

import (
	. "outback/algorithm/common/listnode"
)

func main() {

}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}

	up, down, left, right := 0, m-1, 0, n-1
	cur := head
	for cur != nil {
		for i := left; i <= right; i++ {
			if cur == nil {
				return res
			}
			res[up][i] = cur.Val
			cur = cur.Next
		}
		up++
		for i := up; i <= down; i++ {
			if cur == nil {
				return res
			}
			res[i][right] = cur.Val
			cur = cur.Next
		}
		right--
		for i := right; i >= left; i-- {
			if cur == nil {
				return res
			}
			res[down][i] = cur.Val
			cur = cur.Next
		}
		down--
		for i := down; i >= up; i-- {
			if cur == nil {
				return res
			}
			res[i][left] = cur.Val
			cur = cur.Next
		}
		left++
	}
	return res
}
