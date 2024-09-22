package main

func main() {

}

type NumArray struct {
	tr *SegmentTreeDynamic
}

func Constructor(nums []int) NumArray {
	tr := NewSegmentTreeDynamic()
	return NumArray{tr: tr}
}

func (this *NumArray) Update(index int, val int) {
	this.tr.update(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := this.tr.Query(left, right)
	return sum
}

// 线段树节点结构体
type Node struct {
	left, right *Node // 左右子节点
	val, add    int   // 节点值和懒惰标记
}

// SegmentTreeDynamic 结构体
type SegmentTreeDynamic struct {
	N    int // 定义一个较大的范围值
	root *Node
}

// NewSegmentTreeDynamic 创建一个新的线段树实例
func NewSegmentTreeDynamic() *SegmentTreeDynamic {
	return &SegmentTreeDynamic{root: &Node{}}
}

// update 更新区间 [l, r] 的值
func (s *SegmentTreeDynamic) update(node *Node, start, end, l, r, val int) {
	if l <= start && end <= r {
		node.val += (end - start + 1) * val
		node.add += val
		return
	}
	mid := (start + end) >> 1
	s.pushDown(node, mid-start+1, end-mid)
	if l <= mid {
		s.update(node.left, start, mid, l, r, val)
	}
	if r > mid {
		s.update(node.right, mid+1, end, l, r, val)
	}
	s.pushUp(node)
}

// query 查询区间 [l, r] 的和
func (s *SegmentTreeDynamic) query(node *Node, start, end, l, r int) int {
	if l <= start && end <= r {
		return node.val
	}
	mid := (start + end) >> 1
	ans := 0
	s.pushDown(node, mid-start+1, end-mid)
	if l <= mid {
		ans += s.query(node.left, start, mid, l, r)
	}
	if r > mid {
		ans += s.query(node.right, mid+1, end, l, r)
	}
	return ans
}

// pushUp 更新父节点的值
func (s *SegmentTreeDynamic) pushUp(node *Node) {
	node.val = node.left.val + node.right.val
}

// pushDown 向子节点传递懒惰标记
func (s *SegmentTreeDynamic) pushDown(node *Node, leftNum, rightNum int) {
	if node.left == nil {
		node.left = &Node{}
	}
	if node.right == nil {
		node.right = &Node{}
	}
	if node.add == 0 {
		return
	}
	node.left.val += node.add * leftNum
	node.right.val += node.add * rightNum
	node.left.add += node.add
	node.right.add += node.add
	node.add = 0
}

/*

Node: 定义了线段树的节点结构。
SegmentTreeDynamic: 定义了线段树的结构体，包含根节点。
NewSegmentTreeDynamic: 创建一个新的线段树实例。
update: 更新指定区间的值。
query: 查询指定区间的和。
pushUp: 更新父节点的值。
pushDown: 向子节点传递懒惰标记。
*/
