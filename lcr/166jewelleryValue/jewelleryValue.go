package main

func main() {

}

func jewelleryValue2(frame [][]int) int {
	m, n := len(frame), len(frame[0])
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		a := dfs(i-1, j)
		b := dfs(i, j-1)
		ans := max(a, b) + frame[i][j]
		mem[i][j] = ans
		return ans
	}
	ans := dfs(m-1, n-1)
	return ans
}

func jewelleryValue(frame [][]int) int {
	m, n := len(frame), len(frame[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			f[i][j] = max(f[i-1][j], f[i][j-1]) + frame[i-1][j-1]
		}
	}
	return f[m][n]
}
