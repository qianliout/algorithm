package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxNonOverlapping([]int{-2, 6, 6, 3, 5, 4, 1, 2, 8}, 10))
	fmt.Println(maxNonOverlapping([]int{1, 1, 1, 1, 1}, 2))
}

func maxNonOverlapping(nums []int, target int) int {
	n := len(nums)
	sum := 0
	idx := make(map[int]int)
	idx[0] = -1
	ans := 0
	last := -1
	for i := 0; i < n; i++ {
		sum += nums[i]
		pre, ok := idx[sum-target]
		// 找的是前缀和，前缀和是开区间,所以区区间内的第一个元素是 pre+1,又因为不能是空集，所以是 pre+1>last
		if ok && pre+1 > last {
			ans++
			last = i
		}
		idx[sum] = i // 不对提前算好，有两个原因，1是，因为有负数，会有多个相同的前缀和，2，没有办法判断是否重叠
	}
	return ans
}
