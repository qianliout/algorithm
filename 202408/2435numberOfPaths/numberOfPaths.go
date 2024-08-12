package main

import (
	"math"
)

func main() {

}

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, k+1)
		}
	}
	// 为了避免判断是否越界，可以把下标都加一。此时可以设初始值 f[0][1][0]=1（或者 f[1][0][0]=1）简化一点点代码
	f[0][1][0], f[1][0][0] = 1, 1
	mod := int(math.Pow10(9)) + 7
	for i, row := range grid {
		for j, x := range row {
			for v := 1; v < k; v++ {
				f[i+1][j+1][(v+x)%k] = (f[i+1][j][v] + f[i][j+1][v]) % mod
			}
		}
	}
	return f[m][n][0]
}
