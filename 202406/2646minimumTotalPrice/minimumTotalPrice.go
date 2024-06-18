package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumTotalPrice(1, [][]int{}, []int{2}, [][]int{{0, 0}}))
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	// 标记所以能走到终点的节点，红经过了多少次，注意这里的路径一定是要能走到终点
	var dfs1 func(i, fa, end int) bool
	cnt := make([]int, n)
	dfs1 = func(x, fa, end int) bool {
		if x == end {
			cnt[x]++
			return true
		}
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			if dfs1(y, x, end) {
				// 这里容易出错，这里是 cnt[x]++,不是y, 怎么理解呢：
				// y 能走到终点，那到在终点会被增加一次（上面的判断，上面也只会把终点增加），这里是把中间结点增加
				cnt[x]++
				return true
			}
		}
		return false
	}

	for _, ch := range trips {
		start, end := ch[0], ch[1]
		dfs1(start, -1, end)
	}

	// 算价格,计算的结果是 x 减半时这个子数的价格总和，不减半时这个子树的价格总和（这个子树是包括 x 的）
	var dfs2 func(x, fa int) (int, int)

	dfs2 = func(x, fa int) (int, int) {
		notSub := price[x] * cnt[x]
		sub := (price[x] * cnt[x]) / 2
		// x减半之后，x 下面的其他节点就不能减半了
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			nu, su := dfs2(y, x)
			notSub += min(su, nu) // 如果x 没有减半，那么子树可减可不减
			sub += nu             // 如果 x 减半了，那他的子树只能不减半
		}
		return notSub, sub
	}
	ns, su := dfs2(0, -1)
	return max(ns, su)
}
