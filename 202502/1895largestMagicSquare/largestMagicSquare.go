package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestMagicSquare([][]int{{1, 9, 3, 5, 5, 8, 1, 6, 9}, {4, 1, 1, 6, 8, 3, 5, 7, 6}, {9, 8, 4, 7, 2, 4, 9, 2, 7}, {1, 9, 8, 10, 5, 10, 1, 6, 3}}))
}

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 1; k <= min(m, n); k++ {
				if check(grid, i, j, k) {
					ans = max(ans, k)
				}
			}
		}
	}
	return ans
}

func check(grid [][]int, startI, startJ int, k int) bool {
	m, n := len(grid), len(grid[0])
	if startI+k-1 > min(m, n)-1 || startJ+k-1 > min(m, n)-1 {
		return false
	}
	// 每一行、每一列以及两条对角线的和
	// 检查每一行
	// 1 <= grid[i][j] <= 106
	pre := -1
	for i := startI; i < startI+k; i++ {
		num := 0
		for j := startJ; j < startJ+k; j++ {
			num += grid[i][j]
		}

		if pre != -1 && pre != num {
			return false
		}
		pre = num
	}
	// 每一列
	for j := startJ; j < startJ+k; j++ {
		num := 0
		for i := startI; i < startI+k; i++ {
			num += grid[i][j]
		}
		if pre != -1 && pre != num {
			return false
		}
		pre = num
	}
	// 对角线1
	num := 0
	for i, j := startI, startJ; i < startI+k; i, j = i+1, j+1 {
		num += grid[i][j]
	}

	if pre != -1 && pre != num {
		return false
	}
	// 对角线2
	num = 0
	for i, j := startI, startJ+k-1; i < startI+k; i, j = i-1, j+1 {
		num += grid[i][j]
	}
	if pre != -1 && pre != num {
		return false
	}
	return true
}
