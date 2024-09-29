package main

import (
	"math"
	"sort"
)

func main() {

}

type Element struct {
	value int
	index int
}

func maximumValueSum(board [][]int) int64 {
	m, n := len(board), len(board[0])
	res := make([][]Element, m)

	for i := 0; i < m; i++ {
		res[i] = make([]Element, 0, 3)
		for j := 0; j < n; j++ {
			e := Element{-board[i][j], j}
			res[i] = append(res[i], e)
		}
		sort.Slice(res[i], func(k, l int) bool {
			return res[i][k].value < res[i][l].value
		})
		if len(res[i]) > 3 {
			res[i] = res[i][:3]
		}
	}

	ans := math.MinInt64
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			for k := j + 1; k < m; k++ {
				for a := 0; a < 3; a++ {
					for b := 0; b < 3; b++ {
						for c := 0; c < 3; c++ {
							if res[i][a].index != res[j][b].index && res[k][c].index != res[j][b].index && res[i][a].index != res[k][c].index {
								currentSum := -res[i][a].value - res[j][b].value - res[k][c].value
								if currentSum > ans {
									ans = currentSum
								}
							}
						}
					}
				}
			}
		}
	}
	return int64(ans)
}
