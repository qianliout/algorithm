package main

func main() {

}

func longestCommonSubsequence(text1 string, text2 string) int {
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
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		if text1[i] == text2[j] {
			ans := dfs(i+1, j+1) + 1
			mem[i][j] = ans
			return ans
		}
		ans := max(dfs(i+1, j), dfs(i, j+1))
		mem[i][j] = ans
		return ans
	}
	ans := dfs(0, 0)
	return ans
}
