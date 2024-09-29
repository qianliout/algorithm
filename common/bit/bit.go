package bit

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

func (b *BIT) Update(idx int, x int) {
	delta := x - b.Data[idx]
	i := idx + 1
	for i < len(b.Tree) {
		b.Tree[i] += delta
		i += i & -i
	}
}

func (b *BIT) Pre(i int) int {
	ans := 0
	for i > 0 {
		ans += b.Tree[i]
		// i = i & (i - 1) // 两种写法是一致的
		i -= i & -i
	}
	return ans
}

func (b *BIT) Query(l, r int) int {
	return b.Pre(r+1) - b.Pre(l)
}

func LowBit(n int) int {
	// x 的二进制中，最低位的 1 以及后面所有 0 组成的数
	return n & -n
}
