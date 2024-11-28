package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchMatrix([][]int{{1, 1}}, 1))
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	le, ri := 0, m*n
	for le < ri {
		mid := le + (ri-le)/2
		i, j := mid/n, mid%n
		if i >= 0 && i < m && j >= 0 && j < n && matrix[i][j] >= target {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if le >= 0 && le < m*n && matrix[le/n][le%n] == target {
		return true
	}
	return false
}
