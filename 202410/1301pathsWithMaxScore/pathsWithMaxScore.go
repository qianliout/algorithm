package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pathsWithMaxScore([]string{"E11", "XXX", "11S"}))
}

// 降维是这一个题目的关键
func pathsWithMaxScore(board []string) []int {
	m, n := len(board), len(board[0])
	grid := make([][]byte, m)
	for i := range grid {
		grid[i] = []byte(board[i])
	}
	// g[idx] = 1 : 代表到达起点的路径只有一条，这样我们就有了一个「有效值」可以滚动下去
	// f[idx] = 0 : 代表在起点得分为 0
	g, f := make([]int, m*n+1), make([]int, m*n+1)
	vi := MaxScore{
		Grid: grid,
		M:    m,
		N:    n,
		Inf:  math.MinInt,
		MOD:  int(math.Pow10(9)) + 7,
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			idx := vi.GetIndex(i, j)
			if i == n-1 && j == m-1 {
				f[idx] = 0 // 得分是0
				g[idx] = 1 // 总的路径是1条
				continue
			}
			if grid[i][j] == 'X' {
				f[idx] = vi.Inf
				continue
			}
			value := int(grid[i][j]) - '0'
			if idx == 0 {
				value = 0
			}

			// 计算当前 i j 这个点的最大得分和路径总数，先赋默认值
			u := vi.Inf
			t := 0

			// 向下
			if i+1 < n {
				curU := f[vi.GetIndex(i+1, j)] + value
				curT := g[vi.GetIndex(i+1, j)]
				res := vi.Update(curU, curT, u, t)
				u, t = res.U, res.T
			}
			// 向右
			if j+1 < m {
				curU := f[vi.GetIndex(i, j+1)] + value
				curT := g[vi.GetIndex(i, j+1)]
				res := vi.Update(curU, curT, u, t)
				u, t = res.U, res.T
			}
			if i+1 < n && j+1 < m {
				curU := f[vi.GetIndex(i+1, j+1)] + value
				curT := g[vi.GetIndex(i+1, j+1)]
				res := vi.Update(curU, curT, u, t)
				u, t = res.U, res.T
			}
			f[idx] = u
			if u < 0 {
				f[idx] = vi.Inf
			}
			g[idx] = t
		}
	}
	// 返回
	ans := make([]int, 2)

	if f[0] != vi.Inf {
		ans[0] = f[0]
		ans[1] = g[0]
	}
	return ans
}

// MaxScore 抽象一个结构，用于存储常用数据，避免参数过多
type MaxScore struct {
	Grid [][]byte
	M, N int
	Inf  int
	MOD  int
}

func (vi *MaxScore) GetIndex(i, j int) int {
	return i*vi.M + j
}

type Ans struct {
	U, T int
}

func (vi *MaxScore) Update(curU, curT, u, t int) Ans {
	//  u 表示之前计算的最大分数,t 表示之前分数对应的路径数
	// curU,curT 表示当前的计算结果
	ans := Ans{U: u, T: t}
	if curU > u {
		ans.U = curU
		ans.T = curT
	} else if curU == u && curT != vi.Inf {
		ans.T += curT
	}
	ans.T = ans.T % vi.MOD
	return ans
}

func (vi *MaxScore) ParseIndex(idx int) (int, int) {
	i := idx / vi.M
	j := idx % vi.M
	return i, j
}
