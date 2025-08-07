package main

func main() {

}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i+1][j], f[i][j+1])
			}
		}
	}
	return f[m][n]
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	var dfs func(i, j int) int
	mem := make([][]int, m+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if mem[i][j] >= 0 {
			return mem[i][j]
		}
		if text1[i] == text2[j] {
			a := dfs(i-1, j-1) + 1
			mem[i][j] = a
			return a
		}
		b := max(dfs(i-1, j), dfs(i, j-1))
		mem[i][j] = b
		return b
	}
	return dfs(m-1, n-1)
}
