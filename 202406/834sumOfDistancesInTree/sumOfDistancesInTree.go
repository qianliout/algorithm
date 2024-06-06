package main

func main() {

}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	g := make([][]int, n)

	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	size := make([]int, n)
	for i := range size {
		size[i] = -1
	}
	ans := make([]int, n)
	// 计算以 x 做为根节点的子树大小（包括自已）
	// 同时计算以0做为根节点的距离
	var dfs1 func(x, fa, depth int) int

	dfs1 = func(x, fa, depth int) int {
		if x < 0 || x >= n {
			return 0
		}
		ans[0] += depth

		res := 1 // 自已

		for _, nx := range g[x] {
			if nx != fa {
				res += dfs1(nx, x, depth+1)
			}
		}
		size[x] = res

		return res
	}

	dfs1(0, -1, 0)

	var reroot func(x, fa int)

	reroot = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				ans[y] = ans[x] + n - 2*size[y]
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}
