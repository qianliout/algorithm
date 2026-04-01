package main

func main() {}
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return // 这里不return也不会有错，但是return最好
		}
		if i > n {
			return
		}
		for j := i; j <= n; j++ {
			path = append(path, j)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ans
}
