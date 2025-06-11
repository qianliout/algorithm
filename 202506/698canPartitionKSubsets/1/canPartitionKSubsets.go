package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartitionKSubsets([]int{2, 2, 2, 2, 3, 4, 5}, 4))
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{10, 1, 10, 9, 6, 1, 9, 5, 9, 10, 7, 8, 5, 2, 10, 8}, 11))
}

/*
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
*/

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
	n := len(nums)
	visit := make([]bool, n)

	var dfs func(start, cur, cnt int) bool
	dfs = func(start, cur, cnt int) bool {

		if cnt == k {
			return true
		}
		if cur == target {
			return dfs(0, 0, cnt+1)
		}

		for i := start; i < n; i++ {
			if visit[i] || nums[i]+cur > target {
				continue
			}

			visit[i] = true
			if dfs(i+1, cur+nums[i], cnt) {
				return true
			}
			visit[i] = false
			// 不加这一 步就会超时,为啥呢
			if cur == 0 {
				return false
			}
		}

		return false

	}

	return dfs(0, 0, 0)
}
