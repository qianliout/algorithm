package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 2))
}
func combine1(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		d := k - len(path)
		if len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}
		// 枚举选那个
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j - 1)
			path = path[:len(path)-1]
		}
	}
	dfs(n)
	return ans
}
func combine2(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		// 剪枝
		if n-i+1 < k-len(path) {
			// 剩下的元素都不够凑 k 个了
			return
		}

		if len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}
		// 枚举选那个
		for j := i; j <= n; j++ {
			path = append(path, j)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ans
}
func combine(n int, k int) [][]int {

	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		// 如果先判断i>n 就会出错，要好好理解
		// 下面i加入 path时会递归到下一层才去判断path 是否可以加入答案
		if len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if i > n {
			return
		}
		// 不选
		dfs(i + 1)
		// 选
		path = append(path, i)
		dfs(i + 1)
		path = path[:len(path)-1]
	}
	dfs(1)
	return ans
}
