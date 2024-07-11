package main

func main() {

}

func minReorder(n int, connections [][]int) int {
	g := make([][]pair, n)
	for _, ch := range connections {
		a, b := ch[0], ch[1]
		g[a] = append(g[a], pair{b, 1})
		g[b] = append(g[b], pair{a, 0})
	}
	var dfs func(i, pa int) int

	dfs = func(i, pa int) int {
		res := 0
		for _, nex := range g[i] {
			if nex.b == pa {
				continue
			}
			res += nex.cost + dfs(nex.b, i)
		}
		return res
	}
	return dfs(0, -1)
}

type pair struct {
	b    int
	cost int
}
