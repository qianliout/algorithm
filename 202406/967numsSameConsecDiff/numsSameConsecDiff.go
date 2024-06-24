package main

import (
	"fmt"
)

func main() {
	fmt.Println(numsSameConsecDiff(2, 1))
}

func numsSameConsecDiff(n int, k int) []int {
	ans := make([]int, 0)
	for i := 1; i <= 9; i++ {
		path := []int{i}
		dfs(n, k, path, &ans)
	}
	return ans
}

func dfs(n int, k int, path []int, ans *[]int) {
	if len(path) > n {
		return
	}
	if len(path) == n {
		*ans = append(*ans, gen(path))
		return
	}

	for i := 0; i <= 9; i++ {
		if len(path) > 0 && abs(path[len(path)-1]-i) == k {
			path = append(path, i)
			dfs(n, k, path, ans)
			path = path[:len(path)-1]
		}
	}
}

func gen(path []int) int {
	ans := 0
	for _, ch := range path {
		ans = ans*10 + ch
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
