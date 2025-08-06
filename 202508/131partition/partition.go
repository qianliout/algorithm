package main

func main() {

}

func partition(s string) [][]string {
	ans := make([][]string, 0)
	ss := []byte(s)
	var dfs func(i int, path []string)
	n := len(s)
	dfs = func(i int, path []string) {
		if i >= n {
			ans = append(ans, append([]string{}, path...))
			return
		}
		for j := i; j < n; j++ {
			if check(string(ss[i : j+1])) {
				path = append(path, string(ss[i:j+1]))
				dfs(j+1, path)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0, []string{})
	return ans
}

func check(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
