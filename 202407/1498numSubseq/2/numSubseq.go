package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(numSubseq([]int{2, 3, 3, 4, 6, 7}, 12))
}

func numSubseq(nums []int, target int) int {
	mod := int(math.Pow10(9)) + 7
	sort.Ints(nums)
	n := len(nums)
	cnt := 0
	for i, ch := range nums {
		// ch 做为最小元素
		le, ri := i, n
		for le < ri {
			mid := le + (ri-le+1)>>1
			if mid >= i && mid < n && nums[mid]+ch <= target {
				le = mid
			} else {
				ri = mid - 1
			}
		}

		if le >= i && le < n && ch+nums[le] <= target {
			if le > i {
				// 这里要好好理解，比如 i =0,le=3,[0,1,2,3],此时，0，3是必须选的，1，2两数可选可不选
				cnt = cnt + pow(2, le-i, mod)
			} else if le == i {
				cnt += 1
			}
			cnt = cnt % mod
		}
	}
	return cnt
}

func pow(a, b, m int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a % m
	}
	c := pow(a, b/2, m)
	if b%2 == 1 {
		return (a * c * c) % m
	}
	return c * c % m
}
