package main

import (
	"fmt"

	. "outback/algorithm/common/segtree4"
)

func main() {
	f := Constructor()
	f.Append(2)
	f.AddAll(3)
	f.Append(7)
	f.MultAll(2)
	fmt.Println(f.GetIndex(0))
}

type Fancy struct {
	tr *SegTree
	L  int
}

func Constructor() Fancy {
	n := 2
	data := make([]int, n)
	tree := NewSegTree(data)
	return Fancy{tr: tree, L: 0}
}

func (this *Fancy) Append(val int) {
	this.tr.Update(this.L, val)
	this.L++
}

func (this *Fancy) AddAll(inc int) {
	this.tr.RangeAdd(0, this.L-1, inc)
}

func (this *Fancy) MultAll(m int) {
	this.tr.RangeMul(0, this.L-1, m)
}

func (this *Fancy) GetIndex(idx int) int {
	return this.tr.Query(idx, idx)
}
