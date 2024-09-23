package main

import (
	"fmt"
	"math"
	"math/bits"
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
	n := int(math.Pow10(5)) + 5
	a := make([]int, n)
	tree := NewSegTree(a)
	return Fancy{
		tree: tree,
		sz:   0,
	}
}

func (this *Fancy) Append(val int) {
	this.sz++
	this.tree.Update(1, this.sz, this.sz, &Pair{Add: val, Mul: 1})
}

func (this *Fancy) AddAll(inc int) {
	if this.sz == 0 {
		return
	}
	this.tree.Update(1, 1, this.sz, &Pair{Add: inc, Mul: 1})
}

func (this *Fancy) MultAll(m int) {
	if this.sz == 0 {
		return
	}
	this.tree.Update(1, 1, this.sz, &Pair{Mul: m})
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

type Data struct {
	Value int
}

type Pair struct {
	Add, Mul int
}

func InitTodo() *Pair {
	return &Pair{Mul: 1}
}

type SegTree struct {
	Node []*Node
	MOD  int
}

func NewSegTree(data []int) *SegTree {
	n := len(data)
	t := make([]*Node, 2<<bits.Len(uint(n-1)))
	tree := &SegTree{
		Node: t,
		MOD:  1e9 + 7,
	}
	// 不能动态开数组，还是会超时
	for i := range tree.Node {
		tree.Node[i] = &Node{
			Data: &Data{},
			Todo: &Pair{Mul: 1},
		}
	}
	tree.Build(data, 1, 1, n)

	return tree
}

type Node struct {
	Left  int
	Right int
	Data  *Data
	Todo  *Pair
}

func (s *SegTree) MergeInfo(a, b *Data) *Data {
	return &Data{Value: (a.Value + b.Value) % s.MOD}
}

func (s *SegTree) Do(rootId int, p *Pair) {
	if p == nil {
		panic("fuck panic")
	}
	if p.Mul == 1 && p.Add == 0 {
		return
	}
	node := s.Node[rootId]
	sz := node.Right - node.Left + 1
	if p.Mul != 1 {
		node.Data.Value = (node.Data.Value * p.Mul) % s.MOD
		node.Todo.Add = (node.Todo.Add * p.Mul) % s.MOD
		node.Todo.Mul = (node.Todo.Mul * p.Mul) % s.MOD
	}
	if p.Add != 0 {
		node.Data.Value = (node.Data.Value + sz*p.Add) % s.MOD
		node.Todo.Add = (node.Todo.Add + p.Add) % s.MOD
	}
}

func (s *SegTree) PushDown(rootId int) {
	v := s.Node[rootId].Todo
	if v.Mul == 1 && v.Add == 0 {
		return
	}
	s.Do(rootId<<1, v)
	s.Do(rootId<<1|1, v)
	s.Node[rootId].Todo = InitTodo()
}

func (s *SegTree) Build(a []int, rootId, l, r int) {
	if s.Node[rootId] == nil {
		s.Node[rootId] = &Node{Todo: InitTodo()}
	}
	s.Node[rootId].Left = l
	s.Node[rootId].Right = r
	if l == r {
		// 下标从1开始
		s.Node[rootId].Data = &Data{Value: a[l-1]}
		return
	}

	mid := (l + r) >> 1
	s.Build(a, rootId<<1, l, mid)
	s.Build(a, rootId<<1|1, mid+1, r)
	s.Maintain(rootId)
}

func (s *SegTree) Maintain(rootId int) {
	s.Node[rootId].Data = s.MergeInfo(s.Node[rootId].Data, s.Node[rootId<<1].Data)
}

func (s *SegTree) Update(rootId, l, r int, v *Pair) {
	root := s.Node[rootId]
	if l <= root.Left && root.Right <= r {
		s.Do(rootId, v)
		return
	}
	s.PushDown(rootId)

	mid := (root.Left + root.Right) >> 1
	if l <= mid {
		s.Update(rootId<<1, l, r, v)
	}
	if mid < r {
		s.Update(rootId<<1|1, l, r, v)
	}
	s.Maintain(rootId)
}

func (s *SegTree) Query(rootId int, l, r int) *Data {
	root := s.Node[rootId]
	if l <= root.Left && root.Right <= r {
		return root.Data
	}
	s.PushDown(rootId)
	mid := (root.Left + root.Right) >> 1
	if r <= mid {
		return s.Query(rootId<<1, l, r)
	}
	if l > mid {
		return s.Query(rootId<<1|1, l, r)
	}
	return s.MergeInfo(s.Query(rootId<<1, l, r), s.Query(rootId<<1|1, l, r))
}
