package main

func main() {

}

func maximalNetworkRank(n int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	cnt := make([]int, n)
	for _, ch := range roads {
		a, b := ch[0], ch[1]
		g[a][b] = 1
		g[b][a] = 1
		cnt[a]++
		cnt[b]++
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(ans, cnt[i]+cnt[j]-g[i][j])
		}
	}
	return ans
}
