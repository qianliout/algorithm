package main

func main() {

}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := make([][]pair, 26)
	cnt := make(map[int]bool)

	for i, ch := range equations {
		x, y, v := getInx(ch[0]), getInx(ch[1]), values[i]
		cnt[x] = true
		cnt[y] = true
		g[x] = append(g[x], pair{y, v})
		g[y] = append(g[y], pair{x, 1 / v})
	}
	var dfs func(x, y int, pre float64) float64

	dfs = func(x, y int, pre float64) float64 {
		if !cnt[x] || !cnt[y] {
			return -1
		}
		if x == y {
			return 1
		}
		for _, ch := range g[x] {
			if ch.to == y {
				return pre * ch.v
			}
			a := dfs(y, ch.to, pre*ch.v)
			if a != -1 {
				return a
			}
		}
		return -1
	}
	ans := make([]float64, len(queries))
	for i, ch := range queries {
		ans[i] = dfs(getInx(ch[0]), getInx(ch[1]), 0)
	}
	return ans

}

type pair struct {
	to int
	v  float64
}

func getInx(a string) int {
	b := int(a[0]) - int('a')
	return b
}
