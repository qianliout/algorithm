package main

func main() {

}

func maxOutput(n int, edges [][]int, price []int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	ans := 0
	var dfs func(i, fa int) (int, int)
	dfs = func(i, fa int) (int, int) {
		p := price[i]
		m1, m2 := p, 0
		for _, j := range g[i] {
			if j == fa {
				continue
			}
			s1, s2 := dfs(j, i)
			ans = max(ans, s1+m2, s2+m1)
			m1 = max(m1, s1+p)
			m2 = max(m2, s2+p)
		}
		return m1, m2

	}

	dfs(0, -1)
	return int64(ans)
}
