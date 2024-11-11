package main

func main() {

}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	g := make([][]int, n)
	for i, ch1 := range bombs {
		x, y, d := ch1[0], ch1[1], ch1[2]
		for j, ch2 := range bombs {
			if i == j {
				continue
			}
			x1, y1 := ch2[0], ch2[1]
			if (x-x1)*(x-x1)+(y-y1)*(y-y1) <= d*d {
				g[i] = append(g[i], j) // i能引爆 j
			}
		}
	}
	visit := make([]bool, n)
	var dfs func(i int) int
	dfs = func(x int) int {
		visit[x] = true
		ans := 1
		for _, y := range g[x] {
			if visit[y] {
				continue
			}
			ans += dfs(y)
		}
		return ans
	}
	ans := 0
	for i := 0; i < n; i++ {
		visit = make([]bool, n)
		ans = max(ans, dfs(i))
	}
	return ans
}

type pair struct {
	x, y int
	d    int
}
