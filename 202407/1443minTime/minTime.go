package main

import (
	"fmt"
)

func main() {
	// fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, true, true, false}))
	// fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, false, true, false}))
	// fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, false, false, false, false, false}))
	fmt.Println(minTime(6, [][]int{{0, 1}, {0, 2}, {1, 3}, {3, 4}, {4, 5}}, []bool{false, true, false, false, true, true}))
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	// 标识有苹果的路
	var dfs1 func(i, fa int) bool

	dfs1 = func(i, fa int) bool {
		res := hasApple[i]
		for _, nx := range g[i] {
			if nx == fa {
				continue
			}
			/*
				// 这样写是不对的，因为如如一个点是有苹果，还是得递归到下面的去
				if hasApple[i]{
					return true
				}
			*/

			// 不能这样写，如果 res为真，那么就不会执行后面的 dfs1
			// res = res || dfs1(nx, i)
			next := dfs1(nx, i)
			res = res || next
		}
		hasApple[i] = res
		return res
	}

	dfs1(0, -1)
	var dfs2 func(i, fa int) int

	dfs2 = func(i, fa int) int {
		if !hasApple[i] {
			return 0
		}
		res := 0
		for _, nx := range g[i] {
			if nx == fa {
				continue
			}
			if !hasApple[nx] {
				continue
			}
			res += dfs2(nx, i) + 2 // 为啥是加2呢，因为一来一回
		}
		return res
	}
	res := dfs2(0, n)
	return res
}
