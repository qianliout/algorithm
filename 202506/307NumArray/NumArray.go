package main

func main() {

}

type NumArray struct {
	bit *BIT
}

func Constructor(nums []int) NumArray {
	return NumArray{bit: NewBit(nums)}
}

func (this *NumArray) Update(index int, val int) {
	this.bit.Update(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	ans := this.bit.Pre(right+1) - this.bit.Pre(left)
	return ans
}

type BIT struct {
	N    int
	Data []int
	Tree []int
}

func NewBit(data []int) *BIT {
	n := len(data)
	b := &BIT{
		N:    n,
		Data: make([]int, n),
		Tree: make([]int, n+1),
	}
	for i := 0; i < n; i++ {
		b.Update(i, data[i])
	}
	return b
}

func (b *BIT) Update(i int, va int) {
	delta := va - b.Data[i]
	// b.Data[i] = va
	i = i + 1
	for i <= b.N {
		b.Tree[i] += delta
		i = i + lowbit(i)
	}
}

func (b *BIT) Pre(i int) int {
	ans := 0
	// 容易出错的点，i>0,如果写i>=0就会进入死循环
	for i > 0 {
		ans += b.Tree[i]
		i -= lowbit(i)
	}
	return ans
}

func lowbit(i int) int {
	return i & (-i)
}
