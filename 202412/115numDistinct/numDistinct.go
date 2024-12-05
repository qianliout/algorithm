package main

func main() {

}

func numDistinct(s string, t string) int {

}

func numDistinct1(s string, t string) int {
	m, n := len(s), len(t)

	var dfs func(i, j int) int

	dfs = func(i, j int) int {

		if s[i-1] == t[i-1] {
			return dfs(i-1, j-1)
		}

	}

}
