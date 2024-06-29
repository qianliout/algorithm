package main

import (
	"fmt"
)

func main() {
	fmt.Println(sellingWood(3, 5, [][]int{{1, 4, 2}, {2, 2, 7}, {2, 1, 3}}))
}

func sellingWood(m int, n int, prices [][]int) int64 {
	priceM := make(map[pair]int)
	for _, ch := range prices {
		priceM[pair{h: ch[0], w: ch[1]}] = ch[2]
	}
	var dfs func(h, w int) int
	mem := make([][]int, m+10)
	for i := 0; i < len(mem); i++ {
		mem[i] = make([]int, n+10)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(h, w int) int {
		if h <= 0 || w <= 0 {
			return 0
		}
		if mem[h][w] != -1 {
			return mem[h][w]
		}
		// 不用判断 h，和 w 是否越界的原因是，字典如果没有就默认是0,但是加上了记忆化就最好判断，但是本题目不判断也行，因为不会越界

		ans := priceM[pair{h: h, w: w}]

		// 因为同时只能是横切或竖切，所以不能用双层循环
		for i := 1; i <= h/2; i++ {
			ans = max(ans, dfs(i, w)+dfs(h-i, w))
		}
		for i := 1; i <= w/2; i++ {
			ans = max(ans, dfs(h, i)+dfs(h, w-i))
		}
		mem[h][w] = ans
		return ans
	}

	a := dfs(m, n)
	return int64(a)
}

type pair struct {
	h, w int
}
