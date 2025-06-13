package main

import (
	"fmt"
)

func main() {
	fmt.Println(partition("aab"))
}

func partition(s string) [][]string {
	ans := make([][]string, 0)
	n := len(s)
	var dfs func(i int, path2 []string)

	dfs = func(start int, path1 []string) {
		if start >= n {
			ans = append(ans, append([]string{}, path1...))
			return
		}
		for j := start + 1; j <= n; j++ {
			if check([]byte(s[start:j])) {
				path1 = append(path1, string(s[start:j]))
				dfs(j, path1)
				path1 = path1[:len(path1)-1]
			}
		}
	}
	dfs(0, []string{})
	return ans
}

func check(path []byte) bool {
	i, j := 0, len(path)-1
	for i < j {
		if path[i] != path[j] {
			return false
		}
		i++
		j--
	}
	return true
}
