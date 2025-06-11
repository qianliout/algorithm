package main

func main() {

}

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	ans := dfs(strs, l-1, m, n)
	return ans
}

// 走到了strs[i]处,还剩下 m个1可用，n 个0可用
func dfs(strs []string, i int, m, n int) int {
	if i < 0 {
		if m >= 0 && n >= 0 {
			return 1
		}
		return 0
	}
	c := strs[i]
	if c == "1" {
		yes := dfs(strs, i-1, m-1, n)
		no := dfs(strs, i-1, m, n)
		return yes + no
	}
	yes := dfs(strs, i-1, m, n-1)
	no := dfs(strs, i-1, m, n)
	return yes + no
}
