package main

func main() {

}

func countSubTrees(n int, edges [][]int, labels string) []int {
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	ans := make([]int, n)

	var dfs func(x, fa int) Counter

	dfs = func(x, fa int) Counter {
		c := Counter{}
		for _, y := range g[x] {
			if y != fa {
				nx := dfs(y, x)
				c.Add(nx)
			}
		}
		c[labels[x]]++
		ans[x] = c[labels[x]]
		return c
	}
	dfs(0, -1)
	return ans
}

type Counter map[uint8]int

func (vi Counter) Add(b Counter) {
	for k, v := range b {
		vi[k] += v
	}
}
