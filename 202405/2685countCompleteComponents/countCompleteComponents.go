package main

func main() {

}

func countCompleteComponents(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	visit := make([]bool, n)
	ans := 0
	for i := 0; i < n; i++ {
		if visit[i] {
			continue
		}
		v, e := dfs(g, i, visit)
		if e == v*(v-1) {
			ans++
		}
	}
	return ans
}

// 统计节点数和边数
// 计算时每个节点会计算两次，比如 a---b,遍历到 a 时会计算 a--b 一条边，遍历到 b 时会计算b--a 的一条边
func dfs(g [][]int, x int, visit []bool) (int, int) {
	v := 1
	visit[x] = true

	nex := g[x]
	// nex 里可能已经被遍历过，但是计算边的时候还是加上了的
	e := len(nex)
	for _, ch := range nex {
		if visit[ch] {
			continue
		}
		visit[ch] = true
		i1, i2 := dfs(g, ch, visit)
		v += i1
		e += i2
	}
	return v, e
}
