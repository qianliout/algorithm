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
	f.Append(3)
	f.AddAll(3)
	fmt.Println(f.GetIndex(0)) // 10
	// f.AddAll(3)
	// f.Append(10)
	// f.MultAll(2)
	// fmt.Println(f.GetIndex(0)) // 26
	// fmt.Println(f.GetIndex(1)) // 34
	// fmt.Println(f.GetIndex(2)) // 20
}

type Fancy struct {
	tr *SegTree
	L  int
}

func Constructor() Fancy {
	n := 3
	data := make([]int, n)
	tree := NewSegTree(data)
	return Fancy{tr: tree, L: 0}
}

func (this *Fancy) Append(val int) {
	this.tr.Update(this.L, this.L, Lazy{AddValue: val})
	this.L++
}

func (this *Fancy) AddAll(inc int) {
	this.tr.Update(0, this.L-1, Lazy{AddValue: inc})
}

func (this *Fancy) MultAll(m int) {
	this.tr.Update(0, this.L-1, Lazy{MulValue: m})
}

func (this *Fancy) GetIndex(idx int) int {
	ans := this.tr.Query(idx, idx)
	return ans
}
