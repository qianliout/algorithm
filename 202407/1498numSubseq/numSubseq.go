package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(numSubseq([]int{3, 5, 6, 7}, 9))
}

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	tem := make([]int, n)
	tem[0] = 1
	mod := int(math.Pow10(9)) + 7
	for i := 1; i < n; i++ {
		tem[i] = (tem[i-1] << 1) % mod
	}
	le, ri := 0, n-1
	ans := 0
	// 最大值最小值可以是一个
	for le <= ri {
		if nums[le]+nums[ri] > target {
			ri--
		} else {
			// nums[le]是必须选的，le 到 ri 之间地数可选不可选，le-ri 之间有（ri-le）个数，可选可不选的选法共有2^(ri-le-1)这么多个
			ans = (ans + tem[ri-le]) % mod
			le++
		}
	}
	return ans
}
