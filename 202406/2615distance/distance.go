package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(distance([]int{1, 3, 1, 1, 2}))
}

func distance(nums []int) []int64 {
	group := make(map[int][]int)
	for i, x := range nums {
		group[x] = append(group[x], i)
	}
	n := len(nums)
	ans := make([]int64, n)
	for _, g := range group {
		if len(g) == 1 {
			ans[g[0]] = 0
			continue
		}
		// cnt := minOperations(g)
		cnt := cal(g)
		for i, x := range cnt {
			ans[g[i]] = x
		}
	}
	return ans
}

func cal(nums []int) []int64 {
	// 先算s[0]
	s := 0
	n := len(nums)
	for i := 1; i < n; i++ {
		s += nums[i] - nums[0]
	}
	ans := make([]int64, n)
	ans[0] = int64(s)
	/*
		分组后，对于其中一个组 a，我们先暴力计算出 a[0] 到其它元素的距离之和，设为 s。
		然后计算 a[1]，我们不再暴力计算，而是思考：s 增加了多少？（增量可以是负数）
		设 n 为 a 的长度，从 a[0] 到 a[1]，有 1 个数的距离变大了 a[1]−a[0]，有 n−1个数的距离变小了 a[1]−a[0]
		对于一般的 i,从 a[i-1] 到 a[i]，有 i 个数的距离变大了 a[i]−a[i-1]，有 n−i个数的距离变小了 a[i]−a[i-1]


	*/
	for i := 1; i < n; i++ {
		a := nums[i] - nums[i-1]
		s = s + (i * a) - (n-i)*a
		// s = s + (2*i-n)*a
		ans[i] = int64(s)
	}
	return ans
}

func minOperations(nums []int) []int64 {
	n := len(nums)
	sort.Ints(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	ans := make([]int64, len(nums))
	for i, ch := range nums {
		// 查的右端点
		j := sort.SearchInts(nums, ch)
		ll := (j * ch) - sum[j]
		rr := sum[n] - sum[j] - (n-j)*ch

		ans[i] = int64(ll + rr)
	}
	return ans
}
