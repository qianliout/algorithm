package main

import (
	"fmt"
)

func main() {
	fmt.Println(constructDistancedSequence(5))
}

func constructDistancedSequence(n int) []int {
	// ans := make([]int, 2*n-1)
	path := make([]int, 2*n-1)
	visit := make([]bool, n+1)
	var dfs func(i int) bool

	// 从第一位开始写数字，这样第一次填满就一定是字典序最大的
	dfs = func(pos int) bool {
		if pos >= len(path) {
			return true
		}
		if path[pos] != 0 {
			return dfs(pos + 1)
		}

		for k := n; k >= 1; k-- {
			if visit[k] {
				continue
			}
			if k == 1 {
				visit[k] = true
				path[pos] = k
				if nex := dfs(pos + 1); nex {
					return true
				}
				path[pos] = 0
				visit[k] = false
			} else if k > 1 {
				if pos+k >= len(path) {
					continue
				}
				if path[pos+k] != 0 {
					continue
				}
				visit[k] = true
				path[pos] = k
				path[pos+k] = k

				if nex := dfs(pos + 1); nex {
					// 如果填正确了就不回溯，这样path 里才会有正确的答案
					return true
				}

				path[pos] = 0
				path[pos+k] = 0

				visit[k] = false
			}

		}
		return false
	}
	dfs(0)
	return path
}
