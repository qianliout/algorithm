package main

func main() {

}

func generateParenthesis2(n int) []string {
	ans := make([]string, 0)
	var dfs func(le, ri int, path []byte)

	dfs = func(le, ri int, path []byte) {
		if le == ri && ri == n {
			ans = append(ans, string(path))
			return
		}
		if ri > le {
			return
		}
		if le < n {
			path = append(path, '(')
			dfs(le+1, ri, path)
			path = path[:len(path)-1]
		}
		if ri < le {
			path = append(path, ')')
			dfs(le, ri+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0, []byte{})
	return ans
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	var dfs func(i, le, ri int)
	path := make([]byte, 2*n)
	dfs = func(i, le, ri int) {
		if i == 2*n {
			ans = append(ans, string(path))
			return
		}
		if ri > le {
			return
		}
		if le < n {
			path[i] = '('
			dfs(i+1, le+1, ri)
		}
		if ri < le {
			path[i] = ')'
			dfs(i+1, le, ri+1)
		}
	}
	dfs(0, 0, 0)
	return ans
}
