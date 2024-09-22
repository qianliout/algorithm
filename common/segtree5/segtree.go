package segtree

/*
因此，我们实际上只需要维护数组 nums1 的区间和即可，我们可以通过线段树来实现。
我们定义线段树的每个节点为 Node，每个节点包含如下属性：
    l：节点的左端点，下标从 1 开始。
    r：节点的右端点，下标从 1 开始。
    s：节点的区间和。
    lazy：节点的懒标记。
线段树主要有以下几个操作：
    build(u, l, r)：建立线段树。
    pushdown(u)：下传懒标记。
    pushup(u)：用子节点的信息更新父节点的信息。
    modify(u, l, r)：修改区间和，本题中是反转区间中的每个数，那么区间和 s=r−l+1−s。
    query(u, l, r)：查询区间和。
我们先算出数组 nums2 的所有数之和，记为 s。
*/

type Lazy struct {
	flip bool
}
type Node struct {
	L, R int
	V    int
}

type SegTree struct {
	Data []int
	Node []Node
	Lazy []Lazy
	Cnt1 int
}

func NewSegTree(data []int) *SegTree {
	n := len(data)
	st := &SegTree{
		Data: data,
		Node: make([]Node, n<<2),
		Lazy: make([]Lazy, n<<2),
		Cnt1: 0,
	}
	st.Build()
	return st
}

func (s *SegTree) Build() {
	s.build(1, 1, len(s.Data))
}

func (s *SegTree) build(no int, l, r int) {
	if l > r {
		return
	}
	s.Node[no].L = l
	s.Node[no].R = r
	if l == r {
		s.Node[no].V = s.Data[l-1]
		return
	}
	mid := l + (r-l)/2
	s.build(no>>1, l, mid)
	s.build(no>>1+1, mid+1, r)
	s.pushUp(no)
}

// 用no的两个子节点更新 no 节点
func (s *SegTree) pushUp(no int) {
	s.Node[no].V = s.Node[no<<1].V + s.Node[no<<1+1].V
}

// 更新 no 的两个子节点
func (s *SegTree) pushDown(no int) {
	// if t.tr[u].lazy == 1 {
	// 	mid := (t.tr[u].l + t.tr[u].r) >> 1
	// 	t.tr[u<<1].s = mid - t.tr[u].l + 1 - t.tr[u<<1].s
	// 	t.tr[u<<1].lazy ^= 1
	// 	t.tr[u<<1|1].s = t.tr[u].r - mid - t.tr[u<<1|1].s
	// 	t.tr[u<<1|1].lazy ^= 1
	// 	t.tr[u].lazy ^= 1
	// }

	if s.Lazy[no].flip {

	}

}
