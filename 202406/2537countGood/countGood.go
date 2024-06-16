package main

import (
	"fmt"
)

func main() {
	fmt.Println(countGood([]int{1, 1, 1, 1, 1}, 10))
	fmt.Println(countGood([]int{3, 1, 4, 3, 2, 2, 4}, 2))
}

func countGood1(nums []int, k int) int64 {
	cnt := make(map[int]int)
	le, ans, pair := 0, 0, 0
	for _, num := range nums {
		pair += cnt[num]
		cnt[num]++
		ans += le
		for pair >= k {
			ans++
			cnt[nums[le]]--
			pair -= cnt[nums[le]]
			le++
		}
	}
	return int64(ans)
}

// 一个子数组 arr 如果有 至少 k 对下标 (i, j) 满足 i < j 且 arr[i] == arr[j] ，那么称它是一个 好 子数组。
func countGood2(nums []int, k int) int64 {
	cnt := make(map[int]int)
	le, ans, pair := 0, 0, 0
	for _, num := range nums {
		pair += cnt[num]
		cnt[num]++
		// 找出最后出现的左端点，如果满足条件，那么左端点左边的元素加上左端点本身都可以做左端点的起点
		// 如果去掉左端点，答案没有小于 k，就可以移动左端点
		for pair-cnt[nums[le]] >= k-1 {
			cnt[nums[le]]--
			pair -= cnt[nums[le]]
			le++
		}
		if pair >= k {
			ans += le + 1
		}
	}
	return int64(ans)
}

func countGood(nums []int, k int) int64 {
	cnt := make(map[int]int)

	le, ans, pair, n := 0, 0, 0, len(nums)
	for ri, num := range nums {
		cnt[num]++
		if cnt[num] >= 2 {
			pair += cnt[num] - 1 // 两个数多一对；三个数多两对；四个数多三对
		}
		for pair >= k {
			ans += n - ri // right及其往后的都是好数组
			// 在移出 le 前，把 le 对数对的影响先取消
			if cnt[nums[le]] >= 2 {
				pair -= cnt[nums[le]] - 1 // 为啥会减一呢： 两个数多一对；三个数多两对；四个数多三对
			}
			cnt[nums[le]]--
			le++
		}
	}
	return int64(ans)
}
