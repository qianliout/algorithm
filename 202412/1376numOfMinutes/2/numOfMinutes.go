package main

import (
	"fmt"
)

func main() {
	fmt.Println(numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}))
	fmt.Println(numOfMinutes(7, 6, []int{1, 2, 3, 4, 5, 6, -1}, []int{0, 6, 5, 4, 3, 2, 1}))
}

// 其实就是树的最大深度的变体
// 尝试bfs
// bfs层序遍历会得到错误的结果，为啥呢，因为每一层不会等上一层全传到再进行传递
func numOfMinutes2(n int, headID int, manager []int, informTime []int) int {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		if manager[i] >= 0 {
			g[manager[i]] = append(g[manager[i]], i)
		}
	}
	ans := 0
	q := []int{headID}
	for len(q) > 0 {
		mx := 0
		lev := make([]int, 0)
		for _, no := range q {
			// 结果是错的 ，为啥呢，因为每一层不会等上一层全传到再进行传递
			mx = max(mx, informTime[no])
			lev = append(lev, g[no]...)
		}
		ans += mx
		q = lev
	}
	return ans
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		if manager[i] >= 0 {
			g[manager[i]] = append(g[manager[i]], i)
		}
	}
	ans := 0
	q := []dis{dis{headID, 0}}
	for len(q) > 0 {
		fi := q[0]
		q = q[1:]
		mx := informTime[fi.idx] + fi.ti
		ans = max(ans, mx)
		for _, nx := range g[fi.idx] {
			q = append(q, dis{nx, informTime[fi.idx] + fi.ti})
		}
	}
	return ans
}

type dis struct {
	idx int
	ti  int
}
