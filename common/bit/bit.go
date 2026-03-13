package bit

// BIT 树状数组
// 树状数组是一种可以动态维护序列前缀和的数据结构，支持在 O(logn) 的时间内单点修改、区间查询。
// 树状数组的每个节点维护了原数组中一段连续的元素的和。
// 树状数组的每个节点的下标 i 对应的原数组的下标范围是 [i−lowbit(i)+1,i]，其中 lowbit(i) 表示 i 的二进制表示中最低位的 1 以及后面所有 0 组成的数。
// 例如，lowbit(6)=2，lowbit(10)=2，lowbit(12)=4。
type BIT struct {
	N    int
	Data []int
	Tree []int
}

func NewBIT(data []int) *BIT {
	n := len(data)
	tree := &BIT{N: n}
	// tree.Build1(data)
	tree.Build2(data)
	return tree
}

func (b *BIT) Build1(data []int) {
	b.Data = make([]int, len(data)) // data不先赋值，下面更新时计算差值
	// 最简单的做法是，把 tree[i] 初始化成 0，然后对每个 nums[i]，调用一次 update(i, nums[i])。具体细节见代码。
	b.Tree = make([]int, len(data)+1)
	for i, c := range data {
		b.Update(i, c)
	}
}

func (b *BIT) Build2(data []int) {
	// 其实可以把这些 update 操作合并到一起。从 1 开始枚举 i，把 nums[i−1] 加到 tree[i] 后，tree[i] 就算好了，
	// 直接把 tree[i] 加到下一个关键区间的元素和中，也就是加到 tree[i+lowbit(i)] 中。
	// 下下一个关键区间的元素和由 tree[i+lowbit(i)] 来更新，我们只需要继续往后枚举 i 就行。
	// 注：类似动态规划的「刷表法」。
	b.Data = data // 和Build1的区别就是先赋值
	b.Tree = make([]int, len(data)+1)
	for i, x := range data {
		i++
		b.Tree[i] += x
		next := i + LowBit(i)
		if next <= len(data) {
			b.Tree[next] += b.Tree[i]
		}
	}
}

//	面试时使用这种方式
//
// 更新原 data[idx]=x
func (b *BIT) Update(idx int, x int) {
	delta := x - b.Data[idx]
	b.Data[idx] = x
	i := idx + 1
	for i < len(b.Tree) {
		b.Tree[i] += delta
		i += LowBit(i)
	}
}

// 求前缀和
func (b *BIT) Pre(i int) int {
	ans := 0
	for i > 0 {
		ans += b.Tree[i]
		// i&(i-1) 是把一个数字最右边的1变成 0
		// i = i & (i - 1) // 两种写法是一致的
		// i -= i & -i
		i = i - LowBit(i)
	}
	return ans
}

// 求区间和
func (b *BIT) Query(l, r int) int {
	return b.Pre(r+1) - b.Pre(l)
}

func LowBit(n int) int {
	// x 的二进制中，最低位的 1 以及后面所有 0 组成的数
	return n & -n
}
