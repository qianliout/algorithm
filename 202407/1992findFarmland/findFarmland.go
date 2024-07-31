package main

import (
	"fmt"
)

func main() {
	fmt.Println(findFarmland([][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}))
}

func findFarmland(land [][]int) [][]int {
	// 题目中说了，农场是矩形
	m, n := len(land), len(land[0])
	ans := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 1 {
				a, b := find(land, i, j)
				ans = append(ans, []int{i, j, a, b})
			}
		}
	}
	return ans
}

func find(land [][]int, i, j int) (int, int) {
	m, n := len(land), len(land[0])
	a, b := i, j

	for k := i + 1; k < m; k++ {
		if land[k][j] != 1 {
			break
		}
		a = k
	}

	for k := j + 1; k < n; k++ {
		if land[i][k] != 1 {
			break
		}
		b = k
	}
	for c := i; c <= a; c++ {
		for d := j; d <= b; d++ {
			land[c][d] = -1
		}
	}

	land[i][j] = 1
	return a, b
}
