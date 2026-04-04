package main

import (
	"slices"
	"sort"
)

func main() {

}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)
	used := make([]bool, n)
	path := make([]int, 0)
	ans := make([][]int, 0)
	var dfs func(i, s int)
	dfs = func(i, s int) {
		if s == target {
			ans = append(ans, slices.Clone(path))
			return
		}
		if i >= n || i < 0 || s > target {
			return
		}
		dfs(i+1, s)

		if used[i] {
			return
		}
		// 不太懂这里的剪枝
		if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
			return
		}
		used[i] = true
		path = append(path, candidates[i])
		dfs(i+1, s+candidates[i])
		used[i] = false
		path = path[:len(path)-1]
	}
	dfs(0, 0)
	return ans
}

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。
*/

func combinationSum2_1(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	path := []int{}
	var dfs func(start, remain int)
	dfs = func(start, remain int) {
		if remain == 0 {
			res = append(res, slices.Clone(path))
			return
		}
		for i := start; i < len(candidates); i++ {
			// 层内去重：同一层只用第一个相同数字
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			// 单调剪枝：后面只会更大，直接停止这一层
			if candidates[i] > remain {
				break
			}
			path = append(path, candidates[i])
			dfs(i+1, remain-candidates[i]) // 每个数只能用一次，所以从 i+1 开始
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return res
}
