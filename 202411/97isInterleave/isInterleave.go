package main

func main() {

}

func isInterleave(s1 string, s2 string, s3 string) bool {
	var dfs func(i, j int) bool
	m, n, k := len(s1), len(s2), len(s3)
	if m+n != k {
		return false
	}

	dfs = func(i, j int) bool {
		if i < 0 {
			return s2[:j+1] == s3[:j+1]
		}
		if j < 0 {
			return s1[:i+1] == s3[:i+1]
		}
		ans := false
		if s1[i] == s3[i+j+1] {
			ans = ans || dfs(i-1, j)
		}
		if s2[j] == s3[i+j+1] {
			ans = ans || dfs(i, j-1)
		}
		return ans
	}
	ans := dfs(m-1, n-1)
	return ans
}
