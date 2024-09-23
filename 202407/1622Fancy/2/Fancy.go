package main

import (
	"fmt"

	. "outback/algorithm/common/segtree7"
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
	sz   int
}

func Constructor() Fancy {
	// n := int(math.Pow10(5)) + 5
	n := 4
	a := make([]int, n)
	tree := NewSegTree(a)
	return Fancy{
		tree: tree,
		sz:   0,
	}
}

func (this *Fancy) Append(val int) {
	this.sz++
	this.tree.Update(1, this.sz, this.sz, Pair{Add: val, Mul: 1})
}

func (this *Fancy) AddAll(inc int) {
	if this.sz == 0 {
		return
	}
	this.tree.Update(1, 1, this.sz, Pair{Add: inc, Mul: 1})
}

func (this *Fancy) MultAll(m int) {
	if this.sz == 0 {
		return
	}
	this.tree.Update(1, 1, this.sz, Pair{Mul: m})
}

func (this *Fancy) GetIndex(idx int) int {
	if idx+1 > this.sz {
		return -1
	}
	return this.tree.Query(1, idx+1, idx+1).Value
}

/**
 * Your Fancy object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Append(val);
 * obj.AddAll(inc);
 * obj.MultAll(m);
 * param_4 := obj.GetIndex(idx);
 */
