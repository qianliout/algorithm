package main

import (
	"fmt"
)

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0)
	left, right, up, down := 0, n-1, 0, m-1
	for {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if len(ans) >= m*n {
			break
		}
	}
	return ans[:m*n]
}
