package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumSetSize([]int{1, 2, 3, 4, 5, 6}, []int{2, 3, 2, 3, 2, 3}))
}

func maximumSetSize(nums1, nums2 []int) int {
	// 1，优先移除重复元素
	// 2,还不够就优先移除交集元素
	// 3，再移除剩下的
	set1 := map[int]int{}

	for _, x := range nums1 {
		set1[x]++
	}
	set2 := map[int]int{}

	for _, x := range nums2 {
		set2[x]++
	}
	common := 0
	for x := range set1 {
		if set2[x] > 0 {
			common++
		}
	}

	n1 := len(set1) // nums1中去重后的元素，那么就有 n-n1 个重复的元素
	n2 := len(set2)

	// 这一步就把重复的元素移除了
	ans := n1 + n2 - common

	// 开始移除
	m := len(nums1) / 2

	// 移除了重复元素还不够
	if n1 > m {
		// 这个时时候优先移除交集的元素（因为移除交集中的元素对结果元影响）,此时我们需要移除 n1-m 个元素，如如 n1-m<=common,那么我们就只移除 n1-m 个元素，此时交集的元素没有移除完,
		// 如果 n1-m>common,那么我们最多移除交集的元素（无影响移除），再移除其他的元素（有影响移除）
		// mn表示只移除交集时最多能移除多少个
		mn := min(n1-m, common)
		n1 -= mn
		common -= mn
		// 移除了交集之后还需要移除的个数
		ans -= n1 - m
	}

	if n2 > m {
		// 上面处理集合时已经更新了 common
		n2 -= min(n2-m, common)
		// 因为不会再用到 common，所以可以用不更新 common 了
		ans -= n2 - m
	}

	return ans
}
