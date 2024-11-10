package main

func main() {

}

func minimumCost(n int, edges [][]int, query [][]int) []int {
	g := make([][]pair, n)
	ccAnd := make([]int, 0) //  记录每个连通块的边权的 AND
	ids := make([]int, n)   // 记录每个点所在连通块的编号
	for i := range ids {
		ids[i] = -1 // -1 & n = n
	}

	for _, ch := range edges {
		x, y, w := ch[0], ch[1], ch[2]
		g[x] = append(g[x], pair{y, w})
		g[y] = append(g[y], pair{x, w})
	}

	var dfs func(x int) int
	dfs = func(x int) int {
		a := -1
		ids[x] = len(ccAnd)
		for _, ch := range g[x] {
			y, w := ch.id, ch.va
			a = a & w
			if ids[y] < 0 {
				a = a & dfs(y)
			}
		}
		return a
	}
	for i := 0; i < n; i++ {
		if ids[i] < 0 {
			ccAnd = append(ccAnd, dfs(i))
		}
	}
	ans := make([]int, len(query))
	for i, ch := range query {
		x, y := ch[0], ch[1]
		if ids[x] != ids[y] {
			ans[i] = -1
		} else {
			ans[i] = ccAnd[ids[x]]
		}
	}
	return ans
}

type pair struct {
	id int
	va int
}
