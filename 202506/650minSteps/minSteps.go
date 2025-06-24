package main

func main() {

}

func minSteps2(n int) int {
	inf := n + 5
	var dfs func(cur, cli int) int

	dfs = func(cur, cli int) int {
		if cur == n {
			return 0
		}
		if cur > n {
			return inf
		}
		ans := inf
		if cli > 0 {
			ans = min(ans, dfs(cur+cli, cli)+1)
		}
		if cur != cli {
			ans = min(ans, dfs(cur, cur)+1)
		}
		return ans
	}
	return dfs(1, 0)
}
