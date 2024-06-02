package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1))
	fmt.Println(numberOfPairs([]int{4, 10}, []int{1, 7}, 3))
	fmt.Println(numberOfPairs2([]int{1, 3, 4}, []int{1, 3, 4}, 1))
}

/*
给你两个整数数组 nums1 和 nums2，长度分别为 n 和 m。同时给你一个正整数 k。
如果 nums1[i] 可以被 nums2[j] * k 整除，则称数对 (i, j) 为 优质数对（0 <= i <= n - 1, 0 <= j <= m - 1）。
返回 优质数对 的总数。
*/

// 极致的优化
func numberOfPairs2(nums1 []int, nums2 []int, k int) int64 {
	cnt := make(map[int]int) // 因子的个数
	ans := 0
	for _, ch := range nums1 {
		if ch%k > 0 {
			continue
		}
		ch = ch / k
		for d := 1; d*d <= ch; d++ {
			if ch%d == 0 {
				cnt[d]++
				// 这里一定是 d*d<ch不能是 <=，因为外层循环会统计 d*d==ch 的情况，
				// 比如找12的因子，找到了一个因子2，那么另外一个 12/2=6，也就是说找到了因子2，也就找到了因子6
				// 因此我们只需要枚举到 d*d<=ch 就行
				if d*d < ch {
					cnt[ch/d]++
				}
			}
		}
	}
	for _, ch := range nums2 {
		ans += int(cnt[ch])
	}
	return int64(ans)
}

func numberOfPairs3(nums1 []int, nums2 []int, k int) int64 {
	cnt := make(map[int]int) // 因子的个数
	ans := 0
	for _, ch := range nums1 {
		if ch%k > 0 {
			continue
		}
		ch = ch / k
		for d := 1; d <= ch; d++ {
			if ch%d == 0 {
				cnt[d]++
			}
		}
	}
	for _, ch := range nums2 {
		ans += int(cnt[ch])
	}
	return int64(ans)

}

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	cnt1 := make(map[int]int)
	for _, ch := range nums1 {
		if ch%k != 0 {
			continue
		}
		cnt1[ch/k]++
	}
	// 如果 nums2中没有重复元素，这一步可以不做,直接遍历 nums2
	cnt2 := make(map[int]int)
	for _, ch := range nums2 {
		cnt2[ch]++
	}
	u := slices.Max(nums1) / k
	ans := 0
	for nu, v := range cnt2 {
		s := 0
		for i := nu; i <= u; i = i + nu {
			s += cnt1[i]
		}
		ans += s * v
	}
	return int64(ans)
}
