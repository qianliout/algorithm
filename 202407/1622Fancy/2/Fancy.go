package main

func main() {

}

const (
	_ = iota
	addOperation
	multOperation
)

type Operation struct {
	typeOf int
	value  int
}

type SegmentTree struct {
	tree []Operation
	size int
}

func NewSegmentTree(size int) *SegmentTree {
	return &SegmentTree{
		tree: make([]Operation, 4*size),
		size: size,
	}
}

func (st *SegmentTree) updateLazy(index int) {
	if st.tree[index].typeOf != _ {
		if index < st.size-1 {
			st.tree[2*index+1] = combine(st.tree[2*index+1], st.tree[index])
			st.tree[2*index+2] = combine(st.tree[2*index+2], st.tree[index])
		}
		st.tree[index] = Operation{_}
	}
}

func (st *SegmentTree) updateRange(l, r, value int, typeOf int, index int, tl, tr int) {
	st.updateLazy(index)
	if l > tr || r < tl {
		return
	}
	if l <= tl && tr <= r {
		st.tree[index] = Operation{typeOf, value}
		st.updateLazy(index)
		return
	}
	tmid := (tl + tr) / 2
	st.updateRange(l, r, value, typeOf, 2*index+1, tl, tmid)
	st.updateRange(l, r, value, typeOf, 2*index+2, tmid+1, tr)
}

func combine(a, b Operation) Operation {
	if a.typeOf == _ {
		return b
	}
	if b.typeOf == _ {
		return a
	}
	if a.typeOf == addOperation {
		b.value += a.value
		return b
	} else if a.typeOf == multOperation {
		if b.typeOf == addOperation {
			b.value *= a.value
		} else if b.typeOf == multOperation {
			b.value = (b.value * a.value) % 1000000007
		}
		return b
	}
	return a
}

type Fancy struct {
	Data []int
	Mod  int
	Tree *SegmentTree
}

func Constructor() Fancy {
	return Fancy{
		Data: make([]int, 0),
		Mod:  1000000007,
		Tree: NewSegmentTree(100000), // 假设最大数组长度为100000
	}
}

func (f *Fancy) Append(val int) {
	f.Data = append(f.Data, val)
}

func (f *Fancy) AddAll(inc int) {
	f.Tree.updateRange(0, len(f.Data)-1, inc, addOperation, 0, 0, f.Tree.size-1)
}

func (f *Fancy) MultAll(m int) {
	f.Tree.updateRange(0, len(f.Data)-1, m, multOperation, 0, 0, f.Tree.size-1)
}

func (f *Fancy) GetIndex(idx int) int {
	if idx < 0 || idx >= len(f.Data) {
		return -1
	}
	return f.Data[idx]
}
