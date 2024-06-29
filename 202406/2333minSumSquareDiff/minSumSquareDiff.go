package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1))
}

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	cnt := make([]int, 100009)
	for i := 0; i < len(nums1); i++ {
		cnt[abs(nums1[i]-nums2[i])]++
	}
	k := k1 + k2
	// cnt已经是排好序的了，不以再排序了
	for i := len(cnt) - 1; i > 0; i-- {
		if k <= 0 {
			break
		}
		chang := min(k, cnt[i])
		cnt[i-1] += chang
		k -= chang
		cnt[i] -= chang
	}
	res := 0

	for i, ch := range cnt {
		res += i * i * ch
	}

	return int64(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 参考：https://leetcode.cn/problems/minimum-sum-of-squared-difference/solutions/1658538/javascript-6118-zui-xiao-chai-zhi-ping-f-pp7t/
