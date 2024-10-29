package main

func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges)
	g := make([][]int, n+1)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	gus := make(map[pair]int)
	for _, ch := range guesses {
		p := pair{fa: ch[0], ch: ch[1]}
		gus[p] = 1
	}
	var dfs func(x, fa int)
	cnt0 := 0
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			if gus[pair{x, y}] == 1 {
				cnt0++
			}
			dfs(y, x)
		}
	}
	dfs(0, -1)

	// 树的换根
	ans := 0
	var dfs2 func(x, fa, cnt int)

	dfs2 = func(x, fa, cnt int) {
		if cnt >= k {
			ans++
		}
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			// 换根dp 的精髓，必须找到这个等式
			// 因此，在计算出 cnt0 后，我们可以再次从 0 出发，DFS 这棵树。从节点 x 递归到节点 y 时：
			//    如果有猜测 [x,y]，那么猜对次数减一。
			//    如果有猜测 [y,x]，那么猜对次数加一。
			cnt2 := cnt - gus[pair{fa: x, ch: y}] + gus[pair{fa: y, ch: x}]
			dfs2(y, x, cnt2)
		}
	}

	dfs2(0, -1, cnt0)
	return ans
}

type pair struct {
	fa, ch int
}
