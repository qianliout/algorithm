package main

func main() {

}

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	m, n := len(s), len(t)
	for i < n {
		if j < m && t[i] == s[j] {
			j++
		}
		i++
	}
	return j == m
}
