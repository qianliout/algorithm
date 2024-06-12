package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(earliestSecondToMarkIndices([]int{2, 2, 0}, []int{2, 2, 2, 2, 3, 2, 2, 1}))
	fmt.Println(earliestSecondToMarkIndices([]int{1, 3}, []int{1, 1, 1, 2, 1, 1, 1}))
	fmt.Println(earliestSecondToMarkIndices([]int{0, 1}, []int{2, 2, 2}))
}

/*
给你两个下标从 1 开始的整数数组 nums 和 changeIndices ，数组的长度分别为 n 和 m 。
一开始，nums 中所有下标都是未标记的，你的任务是标记 nums 中 所有 下标。
从第 1 秒到第 m 秒（包括 第 m 秒），对于每一秒 s ，你可以执行以下操作 之一 ：

	选择范围 [1, n] 中的一个下标 i ，并且将 nums[i] 减少 1 。
	如果 nums[changeIndices[s]] 等于 0 ，标记 下标 changeIndices[s] 。
	什么也不做。

请你返回范围 [1, m] 中的一个整数，表示最优操作下，标记 nums 中 所有 下标的 最早秒数 ，如果无法标记所有下标，返回 -1 。
*/
func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	if n > m {
		return -1
	}
	var check func(mx int) bool
	check = func(mx int) bool {
		mx += n
		last := make([]int, n)
		for i := range last {
			last[i] = -1
		}

		for i, v := range changeIndices[:mx] {
			last[v-1] = i // 课程的考试时间，考试时间越晚更好，
		}
		for _, v := range last {
			if v == -1 { // 有课程没有考试时间
				return false
			}
		}
		cnt := 0
		for i, idx := range changeIndices[:mx] {
			idx -= 1
			if last[idx] == -1 {
				return false
			}
			if last[idx] == i {
				if nums[idx] > cnt {
					return false
				}
				cnt -= nums[idx]
			} else {
				cnt++
			}
		}
		return true
	}
	le := n + sort.Search(m+1-n, check)
	if le > m {
		return -1
	}
	return le
}
