package main

import (
	"math"
	"slices"
)

func main() {

}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	f := make([][]int, n+2)

	for i := range f {
		f[i] = make([]int, n+2)
	}
	for i := 0; i < n; i++ {
		f[i][0] = math.MaxInt
		f[i][n+1] = math.MaxInt
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= n; j++ {
			f[i][j] = min(f[i+1][j-1], f[i+1][j], f[i+1][j+1]) + matrix[i][j-1]
		}
	}
	return slices.Min(f[0])
}
