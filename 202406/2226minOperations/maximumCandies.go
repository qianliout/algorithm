package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	// fmt.Println(maximumCandies([]int{5, 8, 6}, 3))
	// fmt.Println(maximumCandies([]int{4, 7, 5}, 4))
	fmt.Println(maximumCandies([]int{1, 2, 1}, 4))
}

// 返回每个小孩可以拿走的 最大糖果数目 (todo 暂时没有能解决)
func maximumCandies1(candies []int, k int64) int {
	mi := slices.Max(candies)
	// le, ri := 1, mi+1
	var check func(t int) bool

	check = func(t int) bool {
		if t <= 0 {
			return false
		}
		cnt := 0
		for _, ch := range candies {
			cnt += ch / t
		}
		return int64(cnt) <= k
	}
	j := sort.Search(mi+1, check)
	if int64(j) < k {
		return -1
	}
	return j
}

// 返回每个小孩可以拿走的 最大糖果数目
func maximumCandies(candies []int, k int64) int {
	mx := slices.Max(candies)
	le, ri := 1, mx+1
	for le < ri {
		mid := le + (ri-le)/2
		// 求不能分得 k 个小孩的最左值，那么 le-1就是结果
		if mid >= 1 && mid < mx+1 && f(candies, mid) < k {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if le < 1 || f(candies, le-1) < k {
		return 0
	}

	return le - 1
}

// 每一堆至少 t个糖果,可以分得多少个至少 t 个糖果的 堆
func f(nums []int, t int) int64 {
	cnt := 0
	t = max(t, 1)
	for _, ch := range nums {
		cnt += ch / t
	}
	return int64(cnt)
}

func semiOrderedPermutation(nums []int) int {
	a, b, n := 0, 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			a = i
		}
		if nums[i] == n {
			b = i
		}
	}
	ans := a + n - b - 1
	if a > b {
		ans -= 1
	}
	return ans

}
