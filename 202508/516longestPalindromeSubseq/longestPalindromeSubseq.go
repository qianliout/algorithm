package main

func main() {

}

func longestPalindromeSubseq(s string) int {
	ss := []byte(s)
	n := len(s)
	var dfs func(i, j int) int
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		ans := 0
		if ss[i] == ss[j] {
			ans = dfs(i+1, j-1) + 2
		} else {
			ans = max(dfs(i+1, j), dfs(i, j-1))
		}
		mem[i][j] = ans
		return ans
	}
	return dfs(0, n-1)
}

func check(ss []byte) bool {
	l, r := 0, len(ss)-1
	for l < r {
		if ss[l] != ss[r] {
			return false
		}
		l++
		r--
	}
	return true
}
