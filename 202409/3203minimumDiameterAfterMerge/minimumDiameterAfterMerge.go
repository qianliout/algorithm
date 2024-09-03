package main

func main() {

}

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	a := diameter(edges1)
	b := diameter(edges2)

	return max(a, b, (a+1)/2+(b+1)/2+1)
}

// 计算一棵树的直径
func diameter(edges [][]int) int {
	n := len(edges)
	g := make([][]int, n+1)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	res := 0
	var dfs func(x, fa int) int
	dfs = func(x, fa int) int {
		maxL := 0
		for _, y := range g[x] {
			if y != fa {
				sub := dfs(y, x) + 1
				res = max(res, maxL+sub)
				maxL = max(maxL, sub)
			}
		}
		return maxL
	}
	dfs(0, -1)
	return res
}
