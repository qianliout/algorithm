package main

import (
	"fmt"
)

func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
}

// 代码太复杂了，可以考虑怎么简化
func orderOfLargestPlusSign(n int, mines2 [][]int) int {
	mines := make([][]int, n)
	for i := range mines {
		mines[i] = make([]int, n)
		for j := range mines[i] {
			mines[i][j] = 1
		}
	}
	for _, ch := range mines2 {
		mines[ch[0]][ch[1]] = 0
	}

	a := make([][]int, n) // 向左
	b := make([][]int, n) // 向右
	c := make([][]int, n) // 向上
	d := make([][]int, n) // 向下
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		b[i] = make([]int, n)
		c[i] = make([]int, n)
		d[i] = make([]int, n)
	}

	// 从 idx=j 的地方向右看，有多少个连续的1
	// 向右
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				a[i][j] = mines[i][j]
			} else if mines[i][j] == 0 {
				a[i][j] = 0
			} else {
				a[i][j] = a[i][j-1] + 1
			}
		}
	}

	// 向左
	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if j == n-1 {
				b[i][j] = mines[i][j]
			} else if mines[i][j] == 0 {
				b[i][j] = 0
			} else {
				b[i][j] = b[i][j+1] + 1
			}
		}
	}
	// 向下
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if i == 0 {
				c[i][j] = mines[i][j]
			} else if mines[i][j] == 0 {
				c[i][j] = 0
			} else {
				c[i][j] = c[i-1][j] + 1
			}
		}
	}

	// 向上
	for j := 0; j < n; j++ {
		for i := n - 1; i >= 0; i-- {
			if i == n-1 {
				d[i][j] = mines[i][j]
			} else if mines[i][j] == 0 {
				d[i][j] = 0
			} else {
				d[i][j] = d[i+1][j] + 1
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, min(a[i][j], b[i][j], c[i][j], d[i][j]))
		}
	}
	return ans
}
