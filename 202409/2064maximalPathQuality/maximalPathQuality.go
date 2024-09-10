package main

func main() {

}

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	g := make([][]edge, n)
	for _, ch := range edges {
		x, y, t := ch[0], ch[1], ch[2]
		g[x] = append(g[x], edge{y, t})
		g[y] = append(g[y], edge{x, t})
	}
	visit := make([]bool, n)
	visit[0] = true

	var dfs func(x, sumTi, sumV int)
	ans := 0
	dfs = func(x, sumTi, sumV int) {
		if x == 0 {
			ans = max(ans, sumV)
			// 可以晚回走
			// not return
		}
		for _, e := range g[x] {
			y, t := e.to, e.ti
			if sumTi+t > maxTime {
				continue
			}
			if visit[y] {
				dfs(y, sumTi+t, sumV)
			} else {
				visit[y] = true
				dfs(y, sumTi+t, sumV+values[y])
				visit[y] = false
			}
		}
	}
	dfs(0, 0, values[0])
	return ans
}

type edge struct {
	to int
	ti int
}
