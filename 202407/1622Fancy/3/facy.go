package main

import (
	"fmt"
	"math"

	. "outback/algorithm/common/segtree2"
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
	Tree *SegTree
	L    int
}

// 不使用懒更新还是会超时
func Constructor() Fancy {
	// n := 100000
	n := 2
	data := make([]int, n)
	mod := int(math.Pow10(9) + 7)
	tree := NewSegTree(data, mod)
	return Fancy{
		Tree: tree,
		L:    0,
	}
}

func (this *Fancy) Append(val int) {
	this.Tree.Update(this.L, val)
	this.L++
}

func (this *Fancy) AddAll(inc int) {
	this.Tree.RangeAdd(0, this.L-1, inc)
}

func (this *Fancy) MultAll(m int) {
	this.Tree.RangeMul(0, this.L-1, m)
}

func (this *Fancy) GetIndex(idx int) int {
	// return this.Tree.Query(idx,idx) // 这样写会有问题，好奇怪
	if idx >= this.L {
		return -1
	}
	return this.Tree.Data[idx]
}
