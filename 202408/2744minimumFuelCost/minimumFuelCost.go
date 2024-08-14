package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt64)
	fmt.Println(len("9223372036854775807"))
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	g := make([][]int, n)
	for _, ch := range roads {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(x, fa int) int
	ans := 0
	dfs = func(x, fa int) int {
		sz := 1
		for _, y := range g[x] {
			if y != fa {
				sz += dfs(y, x)
			}
		}
		if x > 0 {
			// 两种写法都可以
			ans += (sz + seats - 1) / seats
			// ans += (sz-1)/seats + 1
		}
		return sz
	}
	dfs(0, -1)
	return int64(ans)
}
