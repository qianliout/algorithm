package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6))
	fmt.Println(countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11))
	fmt.Println(countFairPairs([]int{0, 0, 0, 0, 0, 0}, 0, 0))
	fmt.Println(sort.SearchInts([]int{1, 2, 2, 2, 2}, 2))
}

func countFairPairs2(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	ans := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		le := lower - nums[i]
		ri := upper - nums[i]
		num1 := nums[i+1:]

		// 找到为真的最小下标
		left := sort.Search(len(num1), func(j int) bool { return num1[j] >= le })
		right := sort.Search(len(num1), func(j int) bool { return num1[j] > ri })
		ans += right - left
	}
	return int64(ans)
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	var result int64
	// 排序后下标的顺序会变化，但是我们是求下标对的总数，对结果是没有影响的
	sort.Ints(nums)
	for i, n := range nums {
		t := nums[i+1:]
		// 找到排序中可以第一个插入的位置
		// 对于向左，如果有多个元素，那么是插入到最前面那个元素之前
		// 比如原数据组是：[1，2，2，2，2] 现在要插入一个元素2，那么返回的值是 1，也就是插入到第0号元素的后面
		left := sort.SearchInts(t, lower-n)
		// 如果想插入最后面的，那么就写 x+1就行，
		right := sort.SearchInts(t, upper-n+1)
		result += int64(right - left)
	}

	return result
}
