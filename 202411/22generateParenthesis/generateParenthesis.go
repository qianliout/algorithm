package main

func main() {

}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	path := make([]byte, 0)
	var dfs func(le, ri int)
	dfs = func(le, ri int) {
		if le == ri && ri == n {
			ans = append(ans, string(path))
			return
		}
		if ri > le {
			return
		}
		if ri > n || le > n {
			return
		}
		if ri < n {
			path = append(path, '(')
			dfs(le+1, ri)
			path = path[:len(path)-1]
		}
		if le > ri {
			path = append(path, ')')
			dfs(le, ri+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}
