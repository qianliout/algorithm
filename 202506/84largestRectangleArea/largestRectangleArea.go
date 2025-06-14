package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println("出栈时计算:", largestRectangleArea2(heights))
	fmt.Println("入栈时计算:", largestRectangleArea3(heights))
	fmt.Println("入栈时计算优化版:", largestRectangleArea5(heights))
}

func largestRectangleArea(heights []int) int {
	st := make([]int, 0)
	// 单调递增
	ans := 0
	for i, num := range heights {
		ans = max(ans, num)
		for len(st) > 0 && heights[st[len(st)-1]] < num {
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
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
			fmt.Println(top, left, right, ans)
		}
		stark = append(stark, i) // 当前元素入栈
	}

	return ans
}

// 入栈时计算的方法：每次入栈时计算以当前元素为最小高度的矩形
func largestRectangleArea3(heights []int) int {
	ans := 0
	stack := make([]int, 0) // 单调递增栈，存储下标

	for i := 0; i < len(heights); i++ {
		// 维护单调递增栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		// 入栈时计算：以当前元素为最小高度，能构成的最大矩形
		// 左边界：栈顶元素的下一个位置（如果栈空则为0）
		leftBound := 0
		if len(stack) > 0 {
			leftBound = stack[len(stack)-1] + 1
		}

		// 右边界：向右扩展到第一个小于当前高度的位置
		rightBound := i
		for rightBound < len(heights) && heights[rightBound] >= heights[i] {
			rightBound++
		}

		// 计算面积
		width := rightBound - leftBound
		area := heights[i] * width
		ans = max(ans, area)

		stack = append(stack, i)
	}

	return ans
}

// 更优雅的入栈时计算方法：预计算每个位置能扩展的宽度
func largestRectangleArea4(heights []int) int {
	n := len(heights)
	ans := 0
	stack := make([]int, 0) // 单调递增栈

	for i := 0; i < n; i++ {
		// 维护单调递增栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		// 入栈时计算：当前元素能向左扩展的最大宽度
		leftBound := -1
		if len(stack) > 0 {
			leftBound = stack[len(stack)-1]
		}

		// 以当前高度为基准，向右扩展到第一个更小的元素
		rightBound := i
		for rightBound < n && heights[rightBound] >= heights[i] {
			rightBound++
		}
		rightBound-- // 回退到最后一个>=heights[i]的位置

		// 计算以heights[i]为高度的矩形面积
		width := rightBound - leftBound
		ans = max(ans, heights[i]*width)

		stack = append(stack, i)
	}

	return ans
}

// 入栈时计算的优化版本：预处理左右边界
func largestRectangleArea5(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	// 预处理：找每个位置左边第一个更小元素的位置
	left := make([]int, n)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 预处理：找每个位置右边第一个更小元素的位置
	right := make([]int, n)
	stack = make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 计算每个位置为高度的最大矩形面积
	ans := 0
	for i := 0; i < n; i++ {
		width := right[i] - left[i] - 1
		ans = max(ans, heights[i]*width)
	}

	return ans
}
