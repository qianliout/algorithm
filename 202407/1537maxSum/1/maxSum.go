package main

import (
	"math"
)

func main() {
}

// 不能提前取余数
func maxSum(nums1 []int, nums2 []int) int {
	var dfs1 func(i int) int // 遍历nums1
	var dfs2 func(i int) int // 遍历nums2
	idx1, idx2 := cal(nums1, nums2)
	n1, n2 := len(nums1), len(nums2)
	mem1, mem2 := make([]int, n1), make([]int, n2)
	for i := 0; i < n1; i++ {
		mem1[i] = -1
	}
	for i := 0; i < n2; i++ {
		mem2[i] = -1
	}
	mod := int(math.Pow10(9)) + 7

	dfs1 = func(i int) int {
		if i < 0 || i >= n1 {
			return 0
		}
		if mem1[i] != -1 {
			return mem1[i]
		}
		res := nums1[i]
		next := dfs1(i + 1)
		if k, ok := idx2[nums1[i]]; ok {
			next = max(next, dfs2(k+1))
		}
		mem1[i] = res + next
		return mem1[i]
	}
	dfs2 = func(i int) int {
		if i < 0 || i >= n2 {
			return 0
		}
		if mem2[i] != -1 {
			return mem2[i]
		}
		res := nums2[i]
		next := dfs2(i + 1)
		if k, ok := idx1[nums2[i]]; ok {
			next = max(next, dfs1(k+1))
		}
		mem2[i] = res + next
		return mem2[i]
	}
	res1 := dfs1(0)
	res2 := dfs2(0)
	return max(res1, res2) % mod
}

// 计算值的下标
func cal(nums1, nums2 []int) (map[int]int, map[int]int) {
	idx1 := make(map[int]int)
	for i, ch := range nums1 {
		idx1[ch] = i
	}
	idx2 := make(map[int]int)
	for i, ch := range nums2 {
		idx2[ch] = i
	}
	return idx1, idx2
}
