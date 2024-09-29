package bit

type BIT struct {
	Data []int
	Tree []int
	N    int
}

func NewBIT(data []int) *BIT {
	b := &BIT{N: len(data)}
	b.Build1(data)
	return b
}

func (b *BIT) Build2(data []int) {
	b.Data = data
	for i, c := range data {
		i++
		b.Tree[i] += c
		next := i + LowBit(i)
		if next <= b.N {
			b.Tree[next] += b.Tree[i]
		}
	}
}

func (b *BIT) Build1(data []int) {
	b.Data = make([]int, b.N)
	for i, ch := range data {
		b.Update(i, ch)
	}
}

func (b *BIT) Update(idx, x int) {
	delta := b.Data[idx] - x
	b.Data[idx] = x
	i := idx + 1
	// 注意这里是小于等于，因为 tree 比data 大1
	for i <= b.N {
		b.Tree[i] += delta
		i = i + LowBit(i)
	}
}

func (b *BIT) Pre(i int) int {
	ans := 0
	for i > 0 {
		ans += b.Tree[i]
		i -= LowBit(i)
	}

	return ans
}

func (b *BIT) Query(l, r int) int {
	return b.Pre(r+1) - b.Pre(l)
}

func LowBit(x int) int {
	return x & -x
}
