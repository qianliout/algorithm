package main

import (
	"fmt"
	"sort"
)

func main() {
	// findExist()
	findNotExist()
}

// 存在
func findExist() {
	nums := []int{0, 1, 2, 3, 3, 3, 4}
	n := len(nums)
	// 如果存在，就是最早出现的位置，如果不存在就是插入的位置，也就是 >x 的最小位置
	// 所以有几个特别点：
	// 1,如果 x 比最大值都大，刚返回 n,如果比最小值都小，刚返回0
	// 如果要找一个存在的数和右端点:sort.SearchInts(nums, x+1) -1,如果保证存在可以这样写
	j1 := sort.SearchInts(nums, 3)
	fmt.Println("存在且左端点：", j1)
	j2 := sort.SearchInts(nums, 3+1)
	fmt.Println("存在且左端点：", j2)

	// 存在，左端点
	// 这个算法的本质是条件为真的最小索引
	j3 := sort.Search(n, func(k int) bool { return nums[k] >= 3 })
	fmt.Println("存在且左端点：", j3)
	// 存在，右端点：>=x+1 得到的数，再减一
	j4 := sort.Search(n, func(k int) bool { return nums[k] >= 3+1 })
	fmt.Println("存在且右端点：", j4-1)
}

// 存在
func findNotExist() {
	nums := []int{0, 1, 2, 2, 2, 4, 4}
	n := len(nums)
	// 如果存在，就是最早出现的位置，如果不存在就是插入的位置，也就是 >x 的最小位置
	// 所以有几个特别点：
	// 1,如果 x 比最大值都大，刚返回 n,如果比最小值都小，刚返回0
	// 如果要找一个存在的数和右端点:sort.SearchInts(nums, x+1) -1,如果保证存在可以这样写
	j1 := sort.SearchInts(nums, 3)
	fmt.Println("不存在且左端点：", j1)
	j2 := sort.SearchInts(nums, 3+1)
	fmt.Println("不存在且左端点：", j2)

	// 存在，左端点
	// 这个算法的本质是条件为真的最小索引
	j3 := sort.Search(n, func(k int) bool { return nums[k] >= 3 })
	fmt.Println("不存在且左端点：", j3)
	// 存在，右端点：>=x+1 得到的数，再减一
	j4 := sort.Search(n, func(k int) bool { return nums[k] >= 3+1 })
	fmt.Println("不存在且右端点：", j4)
}
