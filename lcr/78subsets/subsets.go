package main

import (
	"sort"
)

func main() {

}

// 子集
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	n := len(nums)
	var dfs func(i int, path []int)
	dfs = func(start int, path []int) {
		if len(path) > 0 {
			ans = append(ans, append([]int{}, path...))
		}
		if start < 0 || start >= n {
			return
		}
		for i := start; i < n; i++ {
			// 没有重复元素
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})

	return ans
}

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	n := len(nums)
	var dfs func(i int, path []int)
	used := make([]bool, n)
	sort.Ints(nums)
	dfs = func(i int, path []int) {
		if len(path) > 0 {
			ans = append(ans, append([]int{}, path...))
		}
		if i >= n {
			return
		}
		for j := i; j < n; j++ {
			if j > 0 && nums[j] == nums[j-1] && !used[j-1] {
				continue
			}
			used[j] = true
			path = append(path, nums[j])
			dfs(j+1, path)
			path = path[:len(path)-1]
			used[j] = false
		}
	}
	dfs(0, []int{})
	return ans
}

// 不使用used数组的简洁版本
func subsetsWithDupSimple(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)

	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		// 每个递归层都添加当前路径
		ans = append(ans, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			// 跳过同层重复元素：如果当前元素与前一个相同，且不是本层第一个元素
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{})
	return ans
}

// 带详细注释的版本，帮助理解去重逻辑
func subsetsWithDupDetailed(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums) // 必须先排序，让相同元素相邻

	var dfs func(start int, path []int, depth int)
	dfs = func(start int, path []int, depth int) {
		// 打印当前状态（调试用）
		// fmt.Printf("深度%d: start=%d, path=%v\n", depth, start, path)

		// 每个递归层都添加当前路径作为一个子集
		ans = append(ans, append([]int{}, path...))

		// 在当前层遍历所有可能的选择
		for i := start; i < len(nums); i++ {
			// 关键去重逻辑：
			// i > start 表示不是本层第一个元素
			// nums[i] == nums[i-1] 表示当前元素与前一个相同
			// 这种情况下跳过，避免在同一层选择重复元素
			if i > start && nums[i] == nums[i-1] {
				// fmt.Printf("跳过重复元素: i=%d, nums[i]=%d\n", i, nums[i])
				continue
			}

			// 选择当前元素
			path = append(path, nums[i])
			// 递归到下一层，注意start变成i+1
			dfs(i+1, path, depth+1)
			// 回溯
			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{}, 0)
	return ans
}
