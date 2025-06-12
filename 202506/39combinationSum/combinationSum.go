package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{10, 2, 2, 7, 6, 2, 5}, 8))
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

// 可重复使用
func combinationSum(nums []int, t int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	var dfs func(i int, path []int, sum int)
	// candidates 的所有元素 互不相同 这个条件很重要，说明不需要对结果去重
	dfs = func(start int, path []int, sum int) {
		if sum == t {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if sum > t {
			return
		}
		if start < 0 {
			return
		}
		// 不选
		dfs(start-1, path, sum)
		// 选
		path = append(path, nums[start])
		dfs(start, path, sum+nums[start])
		path = path[:len(path)-1]
	}
	dfs(n-1, []int{}, 0)
	return ans
}

/*
1 <= candidates.length <= 30
2 <= candidates[i] <= 40 //
candidates 的所有元素 互不相同 有这个条件，说明可以不用去重
1 <= target <= 40
*/

// 只使用一次,且有重复元素，
//
//	[2,3] [2,3] 两组的2是一样的值，位置不同，也认为是一个组合
func combinationSum2(nums []int, t int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	n := len(nums)
	var dfs func(i int, path []int, sum int)
	used := make([]bool, n+2)

	dfs = func(start int, path []int, sum int) {
		if sum == t {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if sum > t {
			return
		}
		if start < 0 || start >= n {
			return
		}

		// 不选
		dfs(start+1, path, sum)
		// 选

		// 已选过
		if used[start] {
			return
		}

		if start > 0 && nums[start-1] == nums[start] && !used[start-1] {
			return
		}

		used[start] = true
		path = append(path, nums[start])
		dfs(start+1, path, sum+nums[start])
		path = path[:len(path)-1]
		used[start] = false
	}
	dfs(0, []int{}, 0)
	return ans
}

/*
1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/

/*
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
*/
func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	var dfs func(start, sum int, path []int)
	used := make([]bool, 20)

	dfs = func(start, sum int, path []int) {
		if sum == n && len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if start >= 10 {
			return
		}
		if len(path) > k || sum > n {
			return
		}
		// 不选
		dfs(start+1, sum, path)
		// 	选
		if used[start] {
			return
		}
		used[start] = true
		path = append(path, start)
		dfs(start+1, sum+start, path)
		path = path[:len(path)-1]
		used[start] = false
	}
	dfs(1, 0, []int{})
	return ans
}

/*
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
*/
// 同样的数可以使用多次
// 请注意，顺序不同的序列被视作不同的组合。 [2,1,1],[1,2,1] 是不同的组合
// nums 中的所有元素 互不相同 TODO 如果这里有相同元素呢
func combinationSum4(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := 0
	var dfs func(sum int)
	used := make([]bool, n)
	dfs = func(sum int) {
		if sum == target {
			ans++
			return
		}

		if sum > target {
			return
		}

		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				dfs(sum)
				used[i] = false
			}
			dfs(sum + nums[i])
		}
	}
	dfs(0)
	return ans
}
