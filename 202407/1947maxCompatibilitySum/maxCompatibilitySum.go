package main

import (
	"math/bits"
)

func main() {

}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m, n := len(students), len(students[0])
	// g[i][j] 表示学生 i 和老师 j的兼容值
	g := make([][]int, m)
	for i := range g {
		g[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				if students[i][k] == mentors[j][k] {
					g[i][j]++
				}
			}
		}
	}

	f := make([]int, 1<<m)
	for mask := 1; mask < 1<<m; mask++ {
		c := bits.OnesCount(uint(mask))
		for i := 0; i < m; i++ {
			if mask&(1<<i) != 0 {
				f[mask] = max(f[mask], f[mask^(1<<i)]+g[c-1][i])
			}
		}
	}
	return f[1<<m-1]
}
