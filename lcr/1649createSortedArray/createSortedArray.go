package main

import (
	"slices"
)

func main() {

}

func createSortedArray(instructions []int) int {
	mod := 1_000_000_007
	n := slices.Max(instructions)
	bit := NewBIT(n + 1)
	ans := 0
	for _, ch := range instructions {
		a := bit.query(ch - 1)
		b := bit.query(n) - bit.query(ch)
		ans = (ans + min(a, b)) % mod

		bit.update(ch)
	}
	return ans % mod
}

// 1 <= instructions.length <= 105
// 1 <= instructions[i] <= 105

type BIT struct {
	N    int
	Tree []int
}

func NewBIT(n int) *BIT {
	b := &BIT{
		N: n,
		// 可以理解成小于等于idx 的个数的前缀和
		Tree: make([]int, n+1),
	}
	return b
}

func (b *BIT) update(idx int) {
	idx = idx + 1
	for idx <= b.N {
		b.Tree[idx]++
		idx += lowBit(idx)
	}
}

// 小于等于 idx 的数据的个数
func (b *BIT) query(idx int) int {
	idx = idx + 1
	ans := 0
	for idx > 0 {
		ans += b.Tree[idx]
		idx -= lowBit(idx)
	}
	return ans
}

func lowBit(n int) int {
	return n & (-n)
}
