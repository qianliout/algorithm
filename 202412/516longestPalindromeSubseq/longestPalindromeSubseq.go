package main

func main() {

}

func longestPalindromeSubseq1(s string) int {
	n := len(s)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for le := n - 1; le >= 0; le-- {
		f[le][le] = 1
		for ri := le + 1; ri < n; ri++ {
			if s[le] == s[ri] {
				f[le][ri] = max(f[le][ri], f[le+1][ri-1]+2)
			} else {
				f[le][ri] = max(f[le][ri], f[le+1][ri], f[le][ri-1])
			}
		}
	}
	return f[0][n-1]
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	var dfs func(i, j int) int
	dfs = func(le, ri int) int {
		if le >= n || ri >= n || le < 0 || ri < 0 || le > ri {
			return 0
		}
		if le == ri {
			return 1
		}
		if s[le] == s[ri] {
			return dfs(le+1, ri-1) + 2
		} else {
			return max(dfs(le+1, ri), dfs(le, ri-1))
		}
	}
	ans := dfs(0, n-1)
	return ans
}

// 在递归的归的过程中
// 通过下面的 dfs 知道，需要从 f(i+1) --> f(i),所以必须先算好 i+1
// 同理，j 的变化值是 从j-1-->j  所以 ri 从前到后
