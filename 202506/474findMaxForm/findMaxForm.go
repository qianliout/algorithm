package main

func main() {

}

func findMaxForm1(strs []string, m int, n int) int {
	l := len(strs)
	ans := 0
	dfs(strs, l-1, m, n, &ans, 0)
	return ans
}

// 走到了strs[i]处,还剩下 m个1可用，n 个0可用
// 不加 cache 会超时
func dfs(strs []string, i int, m, n int, mx *int, path int) {
	if i < 0 {
		if m >= 0 && n >= 0 {
			*mx = max(*mx, path)
		}
		return
	}
	no := count(strs[i])
	dfs(strs, i-1, m-no.zero, n-no.one, mx, path+1)
	dfs(strs, i-1, m, n, mx, path)
}

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	nodes := make([]node, l)
	for i := range nodes {
		nodes[i] = count(strs[i])
	}

	f := make([][][]int, l+1)
	for i := range f {
		f[i] = make([][]int, m+1)
		for j := range f[i] {
			f[i][j] = make([]int, n+1)
		}
	}
	// 初值
	// f[i][m][n] = max(f[i-1][m-no.zero][n-no.one],f[i-1][m][n])
	for i := 0; i < l; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				o, z := nodes[i].one, nodes[i].zero
				f[i+1][j][k] = f[i][j][k]
				if j >= z && k >= o {
					f[i+1][j][k] = max(f[i+1][j][k], f[i][j-z][k-o]+1)
				}
			}
		}
	}
	return f[l][m][n]
}

type node struct {
	one, zero int
}

func count(str string) node {
	ans := node{}
	for _, c := range str {
		if c == '1' {
			ans.one++
		} else {
			ans.zero++
		}
	}
	return ans
}
