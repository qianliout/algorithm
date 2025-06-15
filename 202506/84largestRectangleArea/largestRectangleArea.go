package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3, 2, 1, 5, 5, 5, 6, 2, 3, 2, 1, 5, 6, 2, 3}
	fmt.Println("出栈时计算:", largestRectangleArea2(heights))
	fmt.Println("入栈时计算:", largestRectangleArea(heights))
}

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
func largestRectangleArea2(heights []int) int {
	heights = append(heights, 0)
	heights = append(heights, 0)           // 右边加0，确保所有元素都能出栈
	heights = append([]int{0}, heights...) // 左边加0，作为哨兵
	ans := 0
	stark := make([]int, 0) // 单调递增栈，存储下标

	for i := 0; i < len(heights); i++ {
		for len(stark) > 0 && heights[i] < heights[stark[len(stark)-1]] {
			top := heights[stark[len(stark)-1]]     // 当前要计算的高度
			stark = stark[:len(stark)-1]            // 出栈
			left, right := stark[len(stark)-1], i-1 // 左右边界
			ans = max(ans, top*(right-left))        // 计算面积
		}
		stark = append(stark, i) // 当前元素入栈
	}

	return ans
}

// 这种方式更容易理解
func largestRectangleArea3(heights []int) int {
	n := len(heights)
	left := make([]int, n)  // 左边第一个小于当前元素的位置
	right := make([]int, n) // 右边第一个小于当前元素的位置

	// 初始化边界
	for i := 0; i < n; i++ {
		left[i] = -1 // 左边界初始化为-1
		right[i] = n // 右边界初始化为n
	}

	// 找右边第一个更小的元素
	st := make([]int, 0)
	for i, c := range heights {
		for len(st) > 0 && c < heights[st[len(st)-1]] {
			last := st[len(st)-1]
			right[last] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	// 找左边第一个更小的元素
	st = make([]int, 0)
	for i, c := range heights {
		for len(st) > 0 && c < heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1] // 左边第一个更小元素的位置
		}
		st = append(st, i)
	}

	// 计算最大面积
	ans := 0
	for i := 0; i < n; i++ {
		width := right[i] - left[i] - 1
		ans = max(ans, heights[i]*width)
	}
	return ans
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)  // 左边第一个小于当前元素的位置
	right := make([]int, n) // 右边第一个小于当前元素的位置
	for i := 0; i < n; i++ {
		left[i] = -1 // 左边界初始化为-1
		right[i] = n // 右边界初始化为n
	}
	st := make([]int, 0)
	for i, c := range heights {
		for len(st) > 0 && c < heights[st[len(st)-1]] {
			st = st[:len(st)-1] // 出栈
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1] // 左边第一个小于当前元素的位置
		}
		st = append(st, i)
	}
	st = make([]int, 0)
	for i, c := range heights {
		for len(st) > 0 && c < heights[st[len(st)-1]] {
			last := st[len(st)-1]
			right[last] = i
			st = st[:len(st)-1] // 出栈
		}
		st = append(st, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		width := right[i] - left[i] - 1  // 计算宽度
		ans = max(ans, heights[i]*width) // 计算面积
	}
	return ans
}
