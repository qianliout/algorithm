package main

import (
	"sort"
)

func main() {

}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	path := make([]int, n)
	var dfs func(i int)
	used := make([]bool, n)
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
			path[i] = nums[j]
			used[j] = true
			dfs(i + 1)
			used[j] = false
		}
	}

	dfs(0)
	return ans
}
