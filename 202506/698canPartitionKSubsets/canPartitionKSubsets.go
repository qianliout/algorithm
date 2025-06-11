package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets1([]int{2, 2, 2, 2, 3, 4, 5}, 4))
	fmt.Println(canPartitionKSubsets([]int{2, 2, 2, 2, 3, 4, 5}, 4))
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}

/*
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
*/
func canPartitionKSubsets1(nums []int, k int) bool {
	all := 0
	n := len(nums)
	for _, c := range nums {
		all += c
	}
	if all%k != 0 {
		return false
	}
	target := all / k

	used := make([]bool, n)
	sort.Ints(nums)
	var dfs func(i int, cur int)

	dfs = func(start, cur int) {
		if start < 0 || start >= n {
			return
		}
		if k == 0 {
			return
		}
		// 这种写法不对，比如[1,1,2,3]  这里可以用两个1分别和2组合，但是这种况况是不可以的
		if cur == target {
			k--
			dfs(0, 0)
		}

		for i := start; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			dfs(i+1, cur+nums[i])
			used[i] = false
		}
	}
	dfs(0, 0)
	return k == 0
}

func canPartitionKSubsets2(nums []int, k int) bool {
	// 计算总和
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// 如果总和不能被k整除，直接返回false
	if sum%k != 0 {
		return false
	}

	target := sum / k
	n := len(nums)

	// 优化：对数组进行排序（降序），可以更快地剪枝
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	// 如果最大元素大于目标值，无法分割
	if nums[0] > target {
		return false
	}

	// 记录每个数字是否已使用
	used := make([]bool, n)

	// DFS函数：尝试构建k个子集
	var dfs func(k, startIdx, currSum int) bool
	dfs = func(remainingK, startIdx, currSum int) bool {
		// 已找到k个子集
		if remainingK == 0 {
			return true
		}

		// 当前子集已满足目标和，开始构建下一个子集
		if currSum == target {
			return dfs(remainingK-1, 0, 0)
		}

		// 尝试将未使用的数字加入当前子集
		for i := startIdx; i < n; i++ {
			// 剪枝：跳过已使用的数字
			if used[i] {
				continue
			}

			// 剪枝：跳过重复值（如果前一个相同值未被使用）
			// if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			// 	continue
			// }

			// 剪枝：如果加入当前数字会超过目标和，跳过
			// if currSum+nums[i] > target {
			// 	continue
			// }

			// 选择当前数字
			used[i] = true
			dfs(remainingK, i+1, currSum+nums[i])
			// 继续DFS，从下一个位置开始
			if dfs(remainingK, i+1, currSum+nums[i]) {
				return true
			}

			// 回溯
			used[i] = false

			// 剪枝：如果当前和为0且失败，或者当前值等于目标值但失败，后续尝试无意义
			// if currSum == 0 || currSum+nums[i] == target {
			// 	return false
			// }
		}

		return false
	}

	return dfs(k, 0, 0)
}

func canPartitionKSubsets(nums []int, k int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm%k != 0 {
		return false
	}

	target := sm / k
	sort.Ints(nums)
	mem := make(map[int]bool)
	return dfs(nums, k, target, 0, 0, 0, mem)
}

func dfs(nums []int, k, target int, start int, cur int, cnt int, visit map[int]bool) bool {
	if cnt == k {
		return true
	}
	if cur == target {
		return dfs(nums, k, target, 0, 0, cnt+1, visit)
	}

	for i := start; i < len(nums); i++ {
		if visit[i] || nums[i]+cur > target {
			continue
		}

		visit[i] = true
		if dfs(nums, k, target, i+1, cur+nums[i], cnt, visit) {
			return true
		}
		visit[i] = false
		// 不加这一 步就会超时
		if cur == 0 {
			return false
		}
	}

	return false
}
