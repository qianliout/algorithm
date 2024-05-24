package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSumMinProduct([]int{1, 2, 3, 2}))
	fmt.Println(maxSumMinProduct([]int{1, 1, 3, 2, 2, 2, 1, 5, 1, 5}))

	fmt.Println(maxSumMinProduct2([]int{1, 2, 3, 2}))
	fmt.Println(maxSumMinProduct2([]int{1, 1, 3, 2, 2, 2, 1, 5, 1, 5}))

}

// 暴力解法会超时
func maxSumMinProduct2(nums []int) int {
	base := int(math.Pow(10, 9)) + 7
	ans := 0
	sum := make([]int, len(nums)+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	for i, ch := range nums {
		mi := ch
		for j := i; j < len(nums); j++ {
			mi = min(mi, nums[j])
			ans = max(ans, (sum[j+1]-sum[i])*mi)
		}
	}
	return ans % base
}

func maxSumMinProduct(nums []int) int {
	base := int(math.Pow(10, 9)) + 7
	ans := 0
	sum := make([]int, len(nums)+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}

	left := make([]int, len(nums))
	right := make([]int, len(nums))
	for i := range nums {
		left[i] = -1
		right[i] = len(nums)
	}

	st := make([]int, 0)
	for i, ch := range nums {
		for len(st) > 0 && nums[st[len(st)-1]] >= ch {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		}
		st = append(st, i)
	}

	for i, ch := range nums {
		ans = max(ans, ch*(sum[right[i]]-sum[left[i]+1]))
	}
	return ans % base
}
