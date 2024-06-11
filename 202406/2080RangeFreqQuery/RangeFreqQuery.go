package main

import (
	"sort"
)

func main() {

}

type RangeFreqQuery struct {
	Pos map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	pos := make(map[int][]int)
	for i, v := range arr {
		pos[v] = append(pos[v], i)
	}
	return RangeFreqQuery{Pos: pos}

}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	nums := this.Pos[value]
	if len(nums) == 0 {
		return 0
	}
	// 这里 right+1是一个技巧
	// 如果找不到就是插入的位置
	ri := sort.SearchInts(nums, right)

	if ri < len(nums) && nums[ri] == right {
		ri++
	}
	le := sort.SearchInts(nums, left)
	return ri - le

	// 这样写更好
	/*
		ri := sort.SearchInts(nums, right+1)
		le := sort.SearchInts(nums, left)
		return ri - le

	*/
}
