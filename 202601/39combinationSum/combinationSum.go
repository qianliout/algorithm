package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
func combinationSum2(nums []int, t int) [][]int {
	n := len(nums)
	path := make([]int, 0)
	ans := make([][]int, 0)
	var dfs func(i, s int)
	dfs = func(i, s int) {
		if s < 0 {
			return
		}

		if i >= n {
			// 为啥一定要放到 i>=n 里面增加答案,而77题确要放到外面
			if s == 0 {
				ans = append(ans, append([]int{}, path...))
			}
			return
		}
		dfs(i+1, s) // 不选
		if s >= nums[i] {
			path = append(path, nums[i])
			dfs(i, s-nums[i])
			path = path[:len(path)-1]

		}
	}
	dfs(0, t)
	return ans
}

// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的

func combinationSum(nums []int, t int) [][]int {
	n := len(nums)
	path := make([]int, 0)
	ans := make([][]int, 0)
	var dfs func(i, s int)
	dfs = func(i, s int) {
		if s == 0 {
			ans = append(ans, slices.Clone(path))
			return // 这里一定要return 不然就会有重复
		}
		if i >= n || s < 0 {
			return
		}
		dfs(i+1, s)
		path = append(path, nums[i])
		dfs(i, s-nums[i])
		path = path[:len(path)-1]
	}
	dfs(0, t)
	return ans
}
