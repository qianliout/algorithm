package main

import (
	"fmt"

	. "outback/algorithm/common/segtree6"
)

func main() {
	f := Constructor()
	f.Append(2)
	f.AddAll(3)
	f.Append(7)
	f.MultAll(2)
	a := f.GetIndex(0)
	fmt.Println(`ans is 10,but:`, a)
}

type Fancy struct {
	tree *SegTree
	L    int
}

func Constructor() Fancy {
	n := 3
	data := make([]int, n)
	tree := NewSegTree(data)
	return Fancy{tree: tree}
}

func (this *Fancy) Append(val int) {
	this.L++
	this.tree.Update(this.L, this.L, &Todo{Add: val})
}

func (this *Fancy) AddAll(inc int) {
	this.tree.Update(1, this.L, &Todo{Add: inc})
}

func (this *Fancy) MultAll(m int) {
	this.tree.Update(1, this.L, &Todo{Mul: m})
}

func (this *Fancy) GetIndex(idx int) int {
	if idx+1 > this.L {
		return -1
	}
	ans := this.tree.Query(idx+1, idx+1)
	return ans
}
