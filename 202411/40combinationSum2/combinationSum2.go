package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
func combinationSum21(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := make([][]int, 0)
	path := make([]int, 0)
	n := len(candidates)
	used := make([]bool, n)
	var dfs func(start, sum int)

	dfs = func(start, sum int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		// 只有正数
		if sum > target {
			return
		}
		// 枚举选那一个的思想
		for i := start; i < n; i++ {
			// if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
			// 	continue
			// }
			// 小剪枝：同一层相同数值的结点，从第 2 个开始，候选数更少，结果一定发生重复，因此跳过
			// 这种剪枝方法就可以不用 used数组
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			used[i] = true
			path = append(path, candidates[i])
			dfs(i+1, sum+candidates[i])
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(0, 0)
	return ans
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := make([][]int, 0)
	path := make([]int, 0)
	n := len(candidates)
	var dfs func(start, sum int)

	used := make([]bool, n)

	dfs = func(start, sum int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		// 选或不选的思想
		// 只有正数
		if sum > target || start >= n {
			return
		}
		// 不选
		dfs(start+1, sum)

		if start > 0 && candidates[start] == candidates[start-1] && !used[start-1] {
			return
		}

		// 选
		used[start] = true
		path = append(path, candidates[start])
		dfs(start+1, sum+candidates[start])
		used[start] = false
		path = path[:len(path)-1]
	}
	dfs(0, 0)
	return ans
}
