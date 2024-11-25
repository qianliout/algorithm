package segtree

/*
	全部不用指针，根节点从1开始
	支持区间加法，区间乘法
*/

import (
	"math/bits"
)

var initPair = Pair{Add: 0, Mul: 1}

type Data struct {
	Value int
}

type Pair struct {
	Add, Mul int
}

type SegTree struct {
	Node []Node
	MOD  int
}

func NewSegTree(data []int) *SegTree {
	n := len(data)
	t := make([]Node, 2<<bits.Len(uint(n-1)))
	tree := &SegTree{
		Node: t,
		MOD:  1e9 + 7,
	}

	tree.Build(data, 1, 1, n)

	return tree
}

type Node struct {
	Left  int
	Right int
	Data  Data // 使用指针，优化到极致了都还是超时
	Todo  Pair
}

func (s *SegTree) MergeInfo(a, b Data) Data {
	da := Data{}
	da.Value += a.Value
	da.Value += b.Value
	return da
	// return &Data{Value: (a.Value + b.Value) % s.MOD}
}

// 这个方法是最难理解的，从这个方法可以理解 Todo 的含义，如果没有todo,那么在更新时就需要更新全部区间，
// todo记录的是子区间还没有影响到的值，也就是说：s.Head[rootId].Value 需要保持最新,s.Head[rootId].Todo，是 rootId对应的子区间还没有更新的值
// 所以 Do 函数的意义就是，把 rootId 上应该更新的值更新成最新的,并把 本次需要更新子区单的值记录到 todo 中

func (s *SegTree) Do(rootId int, p Pair) {
	if p.Mul == 1 && p.Add == 0 {
		return
	}
	node := s.Node[rootId]
	sz := node.Right - node.Left + 1
	if p.Mul != 1 {
		// 在乘法更新时通常不需要使用 sz（区间大小），因为乘法操作是对节点当前存储的值进行的直接操作，而不是累加操作。对于线段树中的一个节点来说，它存储的值通常是其覆盖范围内的某些聚合信息（如总和、最小值、最大值等）。
		// 在这个特定的函数中，o.d.v 表示的是节点的某种聚合值。当执行乘法更新时，我们只需要将这个聚合值乘以给定的因子 p.mul 即可。这是因为乘法操作不会受到区间长度的影响——无论区间内有多少个元素，每个元素都应当被相同的因子所乘。
		// 例如，如果 o.d.v 存储的是一个区间的元素之和，那么当我们需要将区间内每个元素都乘以某个常数 k 时，我们只需将 o.d.v 乘以 k，而不需要知道具体的元素数量。
		// 因此，在乘法更新时不使用 sz 是合理的。而加法更新时使用 sz 是因为我们需要将区间内每个元素都加上一个常数值，所以需要知道区间的实际长度来正确地更新聚合值
		node.Data.Value = (node.Data.Value * p.Mul) % s.MOD

		// 在乘法更新时更新加法因子是为了保证后续的加法操作能够正确地反映之前的所有操作。具体原因如下：
		// 懒惰传播：线段树中的懒惰传播机制允许我们延迟一些操作，直到真正需要的时候才执行。这意味着在某些节点上可能会累积多个待执行的操作。
		// 组合操作：假设某个节点已经有一个待执行的加法操作 add 和一个待执行的乘法操作 mul。当新的乘法操作 p.mul 应用到这个节点时，我们需要确保之前的加法操作也能够正确地应用。
		// 具体来说：
		// 当前节点的待执行加法操作 o.todo.add 需要乘以新的乘法因子 p.mul，以确保在最终执行加法操作时，结果仍然正确。
		// 例如：
		// 假设当前节点有 o.todo.add = 5 和 o.todo.mul = 2。
		// 新的乘法操作 p.mul = 3 应用到节点上。
		// 更新后的加法操作应该是 o.todo.add = 5 * 3 = 15，这样在最终执行加法操作时，每个元素先乘以 3 再加 15，而不是先加 5 再乘以 3。
		// 因此，在乘法更新时更新加法因子是为了确保所有操作的顺序和效果正确无误。这样可以避免在后续的加法操作中出现错误的结果。
		node.Todo.Add = (node.Todo.Add * p.Mul) % s.MOD

		node.Todo.Mul = (node.Todo.Mul * p.Mul) % s.MOD
	}
	if p.Add != 0 {
		node.Data.Value = (node.Data.Value + sz*p.Add) % s.MOD
		node.Todo.Add = (node.Todo.Add + p.Add) % s.MOD
	}
	// 都是值引用，所以一定要赋值操作
	s.Node[rootId] = node
}

// 懒惰传播，把 rootId 上记录的 todo 更新到子区单中去,并把 rootId 上的todo 设置成初始值

func (s *SegTree) PushDown(rootId int) {
	v := s.Node[rootId].Todo
	if v.Mul == 1 && v.Add == 0 {
		return
	}

	s.Do(rootId<<1, v)
	s.Do(rootId<<1|1, v)

	s.Node[rootId].Todo = initPair
}

func (s *SegTree) Build(a []int, rootId, l, r int) {
	s.Node[rootId].Left = l
	s.Node[rootId].Right = r
	s.Node[rootId].Todo = initPair

	if l == r {
		// 下标从1开始
		s.Node[rootId].Data.Value = a[l-1]
		return
	}

	mid := (l + r) >> 1
	s.Build(a, rootId<<1, l, mid)
	s.Build(a, rootId<<1|1, mid+1, r)
	s.Maintain(rootId)
}

// 通过rootId两个子节点维护父节点的值

func (s *SegTree) Maintain(rootId int) {
	s.Node[rootId].Data = s.MergeInfo(s.Node[rootId<<1].Data, s.Node[rootId<<1|1].Data)
}

// 区单更新

func (s *SegTree) Update(rootId, l, r int, v Pair) {
	root := s.Node[rootId]
	// 全包含了
	if l <= root.Left && root.Right <= r {
		// 把rootId 的值更新成最新的，并把本次更新的值记录到 todo 中
		s.Do(rootId, v)
		return
	}
	// 把todo 记录到两个子节点中去
	s.PushDown(rootId)

	mid := (root.Left + root.Right) >> 1
	if l <= mid {
		s.Update(rootId<<1, l, r, v)
	}
	if mid+1 <= r {
		s.Update(rootId<<1|1, l, r, v)
	}
	s.Maintain(rootId)
}

func (s *SegTree) Query(rootId int, l, r int) Data {
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
	// ans := 0
	// if l <= mid {
	// 	ans += s.Query(rootId<<1, l, r).Value
	// 	// s.Update(rootId<<1, l, r, v)
	// }
	// if mid+1 <= r {
	// 	ans += s.Query(rootId<<1|1, l, r).Value
	// 	// s.Update(rootId<<1|1, l, r, v)
	// }
	// return &Data{Value: ans % s.MOD}
	// // return  ans % s.MOD
}
