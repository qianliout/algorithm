package main

func main() {

}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	var dfs func(left, right int, path []byte)
	dfs = func(i, j int, path []byte) {
		if len(path) == 2*n {
			ans = append(ans, string(path))
			return
		}
		if i < n {
			path = append(path, '(')
			dfs(i+1, j, path)
			path = path[:len(path)-1]
		}
		if j < i {
			path = append(path, ')')
			dfs(i, j+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0, []byte{})
	return ans
}
