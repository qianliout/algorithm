package main

import "fmt"

func main() {
	fmt.Println(tilingRectangle(2, 3))
	fmt.Println(tilingRectangle(5, 8))
	fmt.Println(tilingRectangle(11, 13))
}

func tilingRectangle(n int, m int) int {
	return dfs(n, m)
}

// 贪心的做法不能得到正确的答案
func dfs(n, m int) int {
	if n > m {
		return dfs(m, n)
	}
	if n == 0 || m == 0 {
		return 0
	}
	if n == m {
		return 1
	}
	return 1 + dfs(n, m-n)
}
