package main

func main() {

}

func validArrangement(pairs [][]int) [][]int {
	g := map[int][]int{}
	in := make(map[int]int)
	for _, ch := range pairs {
		s, e := ch[0], ch[1]
		g[s] = append(g[s], e)
		in[e]++ //  入度
	}
	start := 0
	for k, v := range g {
		if len(v) == in[k]+1 {
			start = k
			break
		}
		start = k
	}
	ans := make([][]int, 0)
	m := len(pairs)
	var dfs func(int)
	dfs = func(v int) {
		for len(g[v]) > 0 {
			w := g[v][0]
			g[v] = g[v][1:]
			dfs(w) // 先dfs 再写结果，是这个题目的关键
			ans = append(ans, []int{v, w})
		}
	}
	dfs(start)
	// 题目中已经说了一定有答案，所以可以不用判断
	// 最后需要逆序
	for i := 0; i < m/2; i++ {
		ans[i], ans[m-1-i] = ans[m-1-i], ans[i]
	}
	return ans
}
