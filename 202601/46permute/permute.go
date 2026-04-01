package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)
	used := make([]bool, n)
	dfs = func(i int) {
		if len(path) == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}
			used[j] = true
			path = append(path, nums[j])
			dfs(j)
			used[j] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

func permute2(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)
	used := make([]bool, n)
	dfs = func(i int) {
		if len(path) == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}
			used[j] = true
			path = append(path, nums[j])
			dfs(j)
			used[j] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
