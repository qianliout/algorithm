package main

import (
	"sort"
)

func main() {

}

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	nums := make([]pair, n)
	for i := range nums {
		nums[i] = pair{
			plat: plantTime[i],
			grow: growTime[i],
		}
	}

	// 贪心的思路
	// 两个种子交替种植不会让结果变好，只会一样或变差
	// 按生长天数排序，先播种生长天数最长的会使用结果最优
	sort.Slice(nums, func(i, j int) bool { return nums[i].grow >= nums[j].grow })
	ans := 0
	day := 0
	for _, num := range nums {
		day += num.plat
		ans = max(ans, day+num.grow)
	}
	return ans
}

type pair struct {
	plat, grow int
}

// 证明过程：https://leetcode.cn/problems/earliest-possible-day-of-full-bloom/solutions/1200254/tan-xin-ji-qi-zheng-ming-by-endlesscheng-hfwe/
