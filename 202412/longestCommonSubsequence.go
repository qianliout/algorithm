package main

func main() {

}

func longestCommonSubsequence1(text1 string, text2 string) int {
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
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		a := dfs(i-1, j-1)
		b := dfs(i-1, j)
		c := dfs(i, j-1)
		if text1[i] == text2[j] {
			a = a + 1
		}
		ans := max(a, b, c)
		mem[i][j] = ans
		return ans
	}
	ans := dfs(m-1, n-1)
	return ans
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			a := f[i-1][j-1]
			b := f[i-1][j]
			c := f[i][j-1]
			if text1[i-1] == text2[j-1] {
				a = a + 1
			}
			f[i][j] = max(a, b, c)
		}
	}
	return f[m][n]
}
