package main

import (
	"slices"
)

func main() {

}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	grid := make([][]int, m)

	for i, r := range rowSum {
		grid[i] = make([]int, n)
		for j, c := range colSum {
			x := min(r, c)
			grid[i][j] = x
			r -= x
			colSum[j] -= x
		}
	}

	return grid
}

func specialArray(nums []int) int {
	mx := slices.Max(nums)
	for i := 0; i <= mx; i++ {
		if cnt(nums, i) == i {
			return i
		}
	}
	return -1
}
func cnt(nums []int, x int) int {
	a := 0
	for _, ch := range nums {
		if ch >= x {
			a++
		}
	}
	return a
}
