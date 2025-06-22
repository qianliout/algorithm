package main

func main() {

}

func countPairs(n int, edges [][]int) int64 {
	group := make([]int, 0)
	visit := make([]bool, n)
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(o int) int
	dfs = func(o int) int {
		ans := 1
		visit[o] = true
		for _, j := range g[o] {
			if visit[j] {
				continue
			}
			ans += dfs(j)
		}
		return ans
	}
	for i := 0; i < n; i++ {
		if visit[i] {
			continue
		}
		group = append(group, dfs(i))
	}
	sum := 0
	for _, i := range group {
		sum += i
	}
	ans := 0
	for _, i := range group {
		ans += (sum - i) * i
	}
	return int64(ans / 2)
}
