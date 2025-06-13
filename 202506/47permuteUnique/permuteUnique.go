package main

import "sort"

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	path := make([]int, n)
	sort.Ints(nums)
	used := make([]bool, n)
	var dfs func(i int)

	dfs = func(i int) {
		if i >= n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}
			if j > 0 && nums[j] == nums[j-1] && !used[j-1] {
				continue
			}
			used[j] = true
			path[i] = nums[j]
			dfs(i + 1)
			used[j] = false
		}
	}
	dfs(0)
	return ans
}
