package main

import (
	"fmt"
)

func main() {
	fmt.Println(findValidSplit([]int{4, 7, 8, 15, 3, 5}))
	fmt.Println(findValidSplit([]int{4, 7, 15, 8, 3, 5}))
}

func findValidSplit(nums []int) int {
	n := len(nums)
	// 对于一个质因子 p，设它在数组中的最左和最右的位置为 left 和 right。
	// 那么答案是不能在区间 [left,right) 中的。注意区间右端点可能为答案。
	left := make(map[int]int) // left[p]=i 表示质因数 p 首次出现的下标是 i
	right := make([]int, n)   // right[i] 表示左端点为 i 的区间的右端点的最大值
	help := func(x, i int) {
		if l, ok := left[x]; ok {
			right[l] = max(i, right[l])
		} else {
			left[x] = i // 第一次出现
		}
	}
	for i, ch := range nums {
		for _, pr := range getPrim(ch) {
			help(pr, i)
		}

		// for d := 2; d*d <= ch; d++ {
		// 	if ch%d == 0 {
		// 		help(d, i)
		// 		for ch%d == 0 {
		// 			ch = ch / d
		// 		}
		// 	}
		// }
		//
		// if ch > 1 { // 说明还没有分解完
		// 	help(ch, i)
		// }
	}

	mx := 0
	for l, r := range right {
		if l > mx {
			return mx
		}
		mx = max(mx, r)
	}

	return -1
}

func getPrim(a int) []int {
	ans := make([]int, 0)
	for d := 2; d*d <= a; d++ {
		if a%d == 0 {
			ans = append(ans, d)
			for a%d == 0 {
				a = a / d
			}
		}
	}
	// 这一步特别容易忘记，特别容易出错
	if a > 1 {
		ans = append(ans, a)
	}
	return ans
}
