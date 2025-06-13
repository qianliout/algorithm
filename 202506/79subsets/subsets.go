package main

import (
	"sort"
)

func main() {

}

// 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	n := len(nums)
	used := make([]bool, n)
	var dfs func(i int, path []int)
	dfs = func(start int, path []int) {
		if len(path) > 0 {
			ans = append(ans, append([]int{}, path...))
		}
		if start < 0 || start >= n {
			return
		}
		for i := start; i < n; i++ {
			if used[i] {
				continue
			}
			// 没有重复元素
			used[i] = true
			path = append(path, nums[i])
			dfs(i+1, path)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})

	return ans
}

// 有重复元素的子集
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	n := len(nums)
	var dfs func(start int, path []int)

	used := make([]bool, n)
	dfs = func(start int, path []int) {
		if len(path) > 0 {
			ans = append(ans, append([]int{}, path...))
			// 	 不能直接返回
		}
		if start >= n || start < 0 {
			return
		}

		for j := start; j < n; j++ {
			if used[j] {
				continue
			}
			if j > 0 && nums[j-1] == nums[j] && !used[j-1] {
				continue
			}

			path = append(path, nums[j])
			used[j] = true
			dfs(j+1, path)
			used[j] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})
	return ans
}
