package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {
	fmt.Println(maxBalancedSubsequenceSum([]int{3, 3, 5, 6}))
	fmt.Println(maxBalancedSubsequenceSum([]int{5, -1, -3, 8}))
	fmt.Println(maxBalancedSubsequenceSum([]int{-2, -1}))
}

/*
nums[i]-nums[j] >=i-j
nums[i]-i >= nums[j]-j
把nums[i]-i 定义成一个整体 b
*/
// 不用数状数组的话会超时，结果是对的

// 树状数组模板（维护前缀最大值）
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) int {
	mx := math.MinInt
	for ; i > 0; i &= i - 1 {
		mx = max(mx, f[i])
	}
	return mx
}

func maxBalancedSubsequenceSum(nums []int) int64 {
	// 离散化 nums[i]-i
	b := slices.Clone(nums)
	for i := range b {
		b[i] -= i
	}
	slices.Sort(b)
	b = slices.Compact(b) // 去重

	// 初始化树状数组
	t := make(fenwick, len(b)+1)
	for i := range t {
		t[i] = math.MinInt
	}

	for i, x := range nums {
		j := sort.SearchInts(b, x-i) + 1 // nums[i]-i 离散化后的值（从 1 开始）
		f := max(t.preMax(j), 0) + x
		t.update(j, f)
	}
	return int64(t.preMax(len(b)))
}
