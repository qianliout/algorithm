package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println(unmarkedSumArray([]int{1, 4, 2, 3}, [][]int{{0, 1}}))

}

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	n, m := len(nums), len(queries)
	ids := make([]int, n)
	ids2 := make([]int, n)
	s := 0
	for i, ch := range nums {
		ids[i] = i
		ids2[i] = i
		s += ch
	}
	// 这种写法是对的，
	slices.SortStableFunc(ids, func(i, j int) int { return nums[i] - nums[j] })
	// 这种写法是错的，为啥呢
	sort.SliceStable(ids2, func(i, j int) bool {
		if nums[i] != nums[j] {
			return nums[i]-nums[j] < 0
		}
		return i < j
	})
	fmt.Println(ids)
	fmt.Println(ids2)
	ans := make([]int64, m)
	// 上面用了固定排序，所以这里 j 是全局的
	// 如果不用稳定排序的话，下面的循环中可以让 j 从0 开始，但是这样会  超时
	j := 0
	for qi, p := range queries {
		i, k := p[0], p[1]
		s -= nums[i]
		nums[i] = 0 // 题目中说了全是正数
		for ; j < n && k > 0; j++ {
			i := ids[j]
			if nums[i] > 0 {
				s -= nums[i]
				nums[i] = 0
				k--
			}
		}
		ans[qi] = int64(s)
	}

	return ans
}
