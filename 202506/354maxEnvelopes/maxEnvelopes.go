package main

import (
	"fmt"
	"sort"
)

func main() {
	// 测试用例1
	fmt.Println(maxEnvelopes([][]int{{10, 8}, {1, 12}, {6, 15}, {2, 18}})) // 2

	// 测试用例2
	fmt.Println(maxEnvelopes([][]int{{1, 1}, {2, 2}, {3, 3}})) // 3

	// 测试用例3
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}})) // 3
}

// 不用二分会超时
func maxEnvelopes2(envelopes [][]int) int {
	// 关键修正1：排序策略 - 宽度升序，宽度相同时高度降序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1] // 高度降序！
	})

	n := len(envelopes)
	if n == 0 {
		return 0
	}

	// 关键修正2：使用动态规划，计算以每个位置结尾的最长递增子序列
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 // 每个信封至少可以单独成为一个序列
	}

	maxLen := 1

	// 对每个位置，计算以它结尾的最长递增子序列
	for i := 1; i < n; i++ {
		w, h := envelopes[i][0], envelopes[i][1]
		for j := 0; j < i; j++ {
			w1, h1 := envelopes[j][0], envelopes[j][1]
			// 严格小于条件：w1 < w && h1 < h
			if w1 < w && h1 < h {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

/*
给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
注意：不允许旋转信封。
*/

// O(n log n) 二分查找优化版本
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	// 关键：排序策略 - 宽度升序，宽度相同时高度降序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1] // 高度降序！
	})

	// 提取高度数组，问题转化为求最长递增子序列(LIS)
	heights := make([]int, len(envelopes))
	for i, env := range envelopes {
		heights[i] = env[1]
	}

	// 使用二分查找求LIS
	return lengthOfLIS(heights)
}

// 二分查找优化的LIS算法 - O(n log n)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// tails[i] 表示长度为 i+1 的递增子序列的最小尾部元素
	tails := make([]int, 0)

	for _, num := range nums {
		// 二分查找第一个 >= num 的位置
		left, right := 0, len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 如果 num 比所有元素都大，扩展数组
		if left == len(tails) {
			tails = append(tails, num)
		} else {
			// 否则替换找到的位置，保持tails数组的最优性
			tails[left] = num
		}
	}

	return len(tails)
}
