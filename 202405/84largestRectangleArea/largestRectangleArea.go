package main

func main() {

}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	ans := 0
	// stark 保存的是下标
	// 递增栈
	stark := make([]int, 0)

	for i := 0; i < len(heights); i++ {
		for len(stark) > 0 && heights[i] < heights[stark[len(stark)-1]] {
			top := heights[stark[len(stark)-1]]
			stark = stark[:len(stark)-1]
			left, right := stark[len(stark)-1], i-1
			ans = max(ans, top*(right-left))
		}
		stark = append(stark, i)
	}

	return ans
}
