package main

func main() {}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	path := make([]byte, 0)

	var dfs func(l, r int)

	dfs = func(l, r int) {
		if len(path) == 2*n {
			ans = append(ans, string(path))
			return
		}
		if l < r && r < n {
			path = append(path, ')')
			dfs(l, r+1)
			path = path[:len(path)-1]
		}
		if l < n {
			path = append(path, '(')
			dfs(l+1, r)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}
