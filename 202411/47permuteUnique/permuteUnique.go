package main

import (
	"sort"
)

func main() {

}

func permuteUnique1(nums []int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)
	n := len(nums)
	var dfs func(path []int)
	visit := make([]bool, n)
	dfs = func(path []int) {
		if len(path) == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			// 这种写法也是可以的
			// if i > 0 && nums[i] == nums[i-1] && !visit[i-1] {
			// 	continue
			// }
			if i > 0 && nums[i] == nums[i-1] && visit[i-1] {
				continue
			}
			if visit[i] {
				continue
			}
			visit[i] = true
			path = append(path, nums[i])
			dfs(path)
			path = path[:len(path)-1]
			visit[i] = false
		}
	}
	dfs([]int{})
	return ans
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)
	n := len(nums)
	var dfs func(i int)
	visit := make([]bool, n)
	path := make([]int, n)
	dfs = func(j int) {
		if j >= n {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := 0; i < n; i++ {

			if visit[i] {
				continue
			}
			// 这种写法也是可以的
			// 且效率高些，剪枝彻底些
			// 这种写法容易理解一点
			// 剪枝条件 1：用过的元素不能再使用，
			// 剪枝条件 2：当当前元素和前一个元素值相同（此处隐含这个元素的 index>0 ），并且前一个元素还没有被使用过的时候，我们要剪枝
			// if i > 0 && nums[i] == nums[i-1] && !visit[i-1] {
			// 	continue
			// }

			// 剪枝条件 1：用过的元素不能再使用，
			// 剪枝条件 2：当当前元素和前一个元素值相同（此处隐含这个元素的 index>0 ），并且前一个元素已使用过的时候，我们要剪枝
			if i > 0 && nums[i] == nums[i-1] && visit[i-1] {
				continue
			}

			path[j] = nums[i]

			visit[i] = true
			dfs(j + 1)
			visit[i] = false
		}
	}
	dfs(0)
	return ans
}
