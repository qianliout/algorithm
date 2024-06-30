package main

import (
	"fmt"
)

func main() {
	fmt.Println(arrayChange([]int{1, 2, 4, 6}, [][]int{{1, 3}, {4, 7}, {6, 1}}))
	fmt.Println(arrayChange([]int{1, 2}, [][]int{{1, 3}, {2, 1}, {3, 1}}))
}

// 正着模拟
func arrayChange1(nums []int, operations [][]int) []int {
	ids := make(map[int]int)
	for i, ch := range nums {
		ids[ch] = i
	}

	for _, ch := range operations {
		x, y := ch[0], ch[1]
		id := ids[x]
		nums[id] = y
		delete(ids, x)
		ids[y] = id
	}
	return nums
}

// 倒着遍历
func arrayChange(nums []int, operations [][]int) []int {
	used := make(map[int]int)
	for i := len(operations) - 1; i >= 0; i-- {
		ch := operations[i]
		x, y := ch[0], ch[1]
		if mpY, ok := used[y]; ok {
			y = mpY
		}
		used[x] = y
	}
	for i, ch := range nums {
		if used[ch] > 0 {
			nums[i] = used[ch]
		}
	}

	return nums
}
