package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(waysToSplit([]int{1, 2, 2, 2, 5, 0}))
}

func waysToSplit(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	mod := int(math.Pow10(9)) + 7
	// left<=mid
	// mid<=(all-left)/2
	// 所以对于一个确定的第一分隔点，可能求右边界
	ans := 0
	for i := 1; i < n-1; i++ {
		left := find1(sum, n, i)
		right := find2(sum, n, i)
		if right < left || left >= n-1 || left < i || right < i || right >= n-1 {
			continue
		}
		ans += right - left

	}
	return ans % mod

}

// 求左边，其中 i
func find1(sum []int, n int, a int) int {
	left := sum[a]
	mx := n - 1
	le, ri := a, mx

	for le < ri {
		mid := le + (ri-le)/2
		midValue := sum[mid+1] - sum[a]

		if le < mx && (sum[n]-left)/2 >= midValue && left <= midValue {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}

// 求右边
func find2(sum []int, n int, a int) int {
	left := sum[a]
	mx := n - 1
	le, ri := a, mx

	for le < ri {
		mid := le + (ri-le+1)/2
		midValue := sum[mid+1] - sum[a]

		if le < mx && (sum[n]-left)/2 >= midValue && left <= midValue {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	return le
}
