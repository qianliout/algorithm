package main

import (
	"fmt"

	. "outback/algorithm/common/segtree3"
)

func main() {
	c := Constructor([]int{9, -8})
	c.Update(0, 3)
	fmt.Println(c.SumRange(1, 1))
	fmt.Println(c.SumRange(0, 1))
	c.Update(1, -3)
	fmt.Println(c.SumRange(0, 1))
}

type NumArray struct {
	tr *SegTree
}

func Constructor(nums []int) NumArray {
	tr := NewSegTree(nums)
	return NumArray{tr: tr}
}

func (this *NumArray) Update(index int, val int) {
	this.tr.Update(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := this.tr.Query(left, right)
	return sum
}
