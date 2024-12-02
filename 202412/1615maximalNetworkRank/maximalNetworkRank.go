package main

func main() {

}

func maximalNetworkRank(n int, roads [][]int) int {
	in := make([]int, n)
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	for _, ch := range roads {
		x, y := ch[0], ch[1]
		in[x]++
		in[y]++
		g[x][y]++
		g[y][x]++
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(ans, in[i]+in[j]-g[i][j])
		}
	}

	return ans
}
