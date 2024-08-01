package main

import (
	"fmt"
)

func main() {
	fmt.Println(platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}}))
}

func platesBetweenCandles(s string, queries [][]int) []int {
	n, m := len(s), len(queries)
	sum := make([]int, n+1)
	pivotLeft := make([]int, n)  // 左边最近位置的蜡烛
	pivotRight := make([]int, n) // 右边最近位置的蜡烛
	pivotLeft[0] = -1
	pivotRight[n-1] = n

	for i, ch := range s {
		// 左边位置的蜡烛
		if i > 0 {
			pivotLeft[i] = pivotLeft[i-1]
		}
		if ch == '|' {
			pivotLeft[i] = i
		}

		sum[i+1] = sum[i]
		if ch == '*' {
			sum[i+1] += 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i < n-1 {
			pivotRight[i] = pivotRight[i+1]
		}
		if s[i] == '|' {
			pivotRight[i] = i
		}
	}

	ans := make([]int, m)

	for i, ch := range queries {
		ans[i] = 0 // 初值是0
		start, end := pivotRight[ch[0]], pivotLeft[ch[1]]
		if start < end && start != -1 && end != n {
			ans[i] = sum[end+1] - sum[start]
		}
	}
	return ans
}
