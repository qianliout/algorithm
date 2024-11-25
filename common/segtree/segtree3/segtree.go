package segtree

import "math/bits"

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
	// 可以先把数组赋值，下面就可以不用空指针判断
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
	Data  *Data // 使用指针，优化到极致了都还是超时
	Todo  *Pair
}

func (s *SegTree) MergeInfo(a, b *Data) *Data {

	da := &Data{}
	if a != nil {
		da.Value += a.Value
	}
	if b != nil {
		da.Value += b.Value
	}
	return da
	// 如果上面在初始化时已经赋初值了,就不用空指针判断
	// return &Data{Value: (a.Value + b.Value) % s.MOD}
}

func (s *SegTree) Do(rootId int, p *Pair) {
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
	// 因为传的是指针，所以最后可以不用再赋值
	// s.Head[rootId] = node
}

func (s *SegTree) PushDown(rootId int) {
	v := s.Node[rootId].Todo
	if v.Mul == 1 && v.Add == 0 {
		return
	}
	s.Do(rootId<<1, v)
	s.Do(rootId<<1|1, v)
	s.Node[rootId].Todo.Add = 0
	s.Node[rootId].Todo.Mul = 1
	// s.Head[rootId].Todo = InitTodo()
}

func (s *SegTree) Build(a []int, rootId, l, r int) {
	if s.Node[rootId] == nil {
		s.Node[rootId] = &Node{Data: &Data{}, Todo: InitTodo()}
	}
	s.Node[rootId].Left = l
	s.Node[rootId].Right = r
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

func (s *SegTree) Maintain(rootId int) {
	s.Node[rootId].Data = s.MergeInfo(s.Node[rootId<<1].Data, s.Node[rootId<<1|1].Data)
}

func (s *SegTree) Update(rootId, l, r int, v *Pair) {
	if s.Node[rootId] == nil {
		s.Node[rootId] = &Node{Data: &Data{}, Todo: InitTodo()}
	}
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
	if mid+1 <= r {
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

	// 这样写也是可以的
	// if r <= mid {
	// 	return s.Query(rootId<<1, l, r)
	// }
	// if l > mid {
	// 	return s.Query(rootId<<1|1, l, r)
	// }
	// return s.MergeInfo(s.Query(rootId<<1, l, r), s.Query(rootId<<1|1, l, r))

	// 更好的写法是这样
	ans := 0
	if l <= mid {
		ans += s.Query(rootId<<1, l, r).Value
	}
	if mid+1 <= r {
		ans += s.Query(rootId<<1|1, l, r).Value
	}
	return &Data{Value: ans % s.MOD}
}
