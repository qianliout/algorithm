package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(recoverArray([]int{2, 10, 6, 4, 8, 12}))
}

func recoverArray(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		d := nums[i] - nums[0]
		if d&1 == 1 {
			continue
		}
		d = d / 2
		ans := helper(nums, d, i)
		if len(ans) > 0 {
			return ans
		}
	}
	return []int{}
}

func helper(nums []int, k, i int) []int {
	n := len(nums)
	visit := make([]bool, n)
	visit[0] = true
	visit[i] = true

	ans := make([]int, 0)

	ans = append(ans, (nums[0]+nums[i])/2) // 第一个数

	for lo, hi := 0, i+1; hi < n; hi++ {
		for lo = lo + 1; visit[lo]; lo++ {
		}
		for ; hi < n && nums[hi]-nums[lo] < 2*k; hi++ {
		}
		if hi == n || nums[hi]-nums[lo] > 2*k {
			break
		}
		visit[hi] = true
		ans = append(ans, (nums[lo]+nums[hi])/2)
	}
	if len(ans) == n/2 {
		return ans
	}
	return []int{}
}
