package main

import (
	"fmt"
)

func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
}

func orderOfLargestPlusSign(n int, mines2 [][]int) int {
	mines := make([][]int, n+10)
	for i := range mines {
		mines[i] = make([]int, n+10)
	}
	for _, ch := range mines2 {
		mines[ch[0]+1][ch[1]+1] = 1
	}

	a := make([][]int, n+10)
	b := make([][]int, n+10) // 向右
	c := make([][]int, n+10) // 向上
	d := make([][]int, n+10) // 向下
	for i := 0; i < n+10; i++ {
		a[i] = make([]int, n+10)
		b[i] = make([]int, n+10)
		c[i] = make([]int, n+10)
		d[i] = make([]int, n+10)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if mines[i][j] == 1 {
				a[i][j] = a[i-1][j] + 1
				b[i][j] = b[i][j-1] + 1
			}

			if mines[n+1-i][n+1-j] == 1 {
				c[n+1-i][n+1-j] = c[n+2-i][n+1-j] + 1
				d[n+1-i][n+1-j] = d[n+1-i][n+2-j] + 1
			}

		}
	}
	// for i := 1; i <= n; i++ {
	// 	for j := n - 1; j >= 0; j-- {
	// 		if mines[i][j] == 1 {
	// 			c[i][j] = c[i+1][j] + 1
	// 			d[i][j] = c[i][j+1] + 1
	// 		}
	// 	}
	// }

	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			ans = max(ans, min(a[i][j], b[i][j], c[i][j], d[i][j]))
		}
	}
	return ans
}
