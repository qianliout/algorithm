package main

import (
	"math"
	"math/bits"
	"sort"
)

func main() {

}

func minChanges(n int, k int) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == j {
			return 0
		}
		if j&1 == 1 && i&1 == 0 {
			return math.MinInt64
		}

		if j&1 == 0 && i&1 == 1 {
			return dfs(i>>1, j>>1) + 1
		}
		return dfs(i>>1, j>>1)
	}
	ans := dfs(n, k)
	if ans < 0 {
		return -1
	}
	return ans
}

/*
给你两个正整数 n 和 k。
你可以选择 n 的 二进制表示 中任意一个值为 1 的位，并将其改为 0。
返回使得 n 等于 k 所需要的更改次数。如果无法实现，返回 -1。
*/

/*
给你一个整数数组 arr 。请你将数组中的元素按照其二进制表示中数字 1 的数目升序排序。
如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列。
请你返回排序后的数组
*/
func sortByBits2(arr []int) []int {
	n := len(arr)
	nums := make([]int, n)
	for i := range nums {
		nums[i] = count(arr[i])
	}
	// 千万不能这样做，这样是错的
	sort.Slice(arr, func(i, j int) bool {
		if nums[i] != nums[j] {
			return nums[i] < nums[j]
		}
		return arr[i] < arr[j]
	})
	return arr
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a := bits.OnesCount(uint(arr[i]))
		b := bits.OnesCount(uint(arr[j]))
		if a < b {
			return true
		} else if a > b {
			return false
		} else {
			return arr[i] < arr[j]
		}
	})
	return arr
}

// 0 <= arr[i] <= 10^4
func count(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return count(n>>1) + n&1
}

func count2(n int) int {
	ans := 0
	for n > 0 {
		ans += n & 1
		n = n >> 1
	}
	return ans
}

func count3(n int) int {
	return bits.OnesCount(uint(n))
}
