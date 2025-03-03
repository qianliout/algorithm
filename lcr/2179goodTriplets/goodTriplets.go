package main

import (
	"fmt"
)

func main() {
	// fmt.Println(goodTriplets([]int{2, 0, 1, 3}, []int{0, 1, 2, 3}))
	fmt.Println(goodTriplets([]int{4, 0, 1, 3, 2}, []int{4, 1, 0, 2, 3}))
	// 输入：nums1 = [4,0,1,3,2], nums2 = [4,1,0,2,3]
	// 输出：4
	// 解释：总共有 4 个好三元组 (4,0,3) ，(4,0,2) ，(4,1,3) 和 (4,1,2) 。
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	p := make([]int, n+1)
	// 这么做是让 bit不要用0开始 从1开始
	// for i := 0; i < n; i++ {
	// 	nums1[i] = nums1[i] + 1
	// 	nums2[i] = nums2[i] + 1
	// }

	for i, ch := range nums1 {
		p[ch] = i
	}
	bit := NewBIT(n + 2)
	ans := 0
	for i := 1; i < n; i++ {
		bit.update(p[nums2[i-1]])
		y := p[nums2[i]]
		less := bit.query(y)

		ans += less * (n - 1 - y - (i - less))
	}

	return int64(ans)
}

type BIT struct {
	N    int
	Tree []int
}

func NewBIT(n int) *BIT {
	b := &BIT{
		N:    n,
		Tree: make([]int, n+1),
	}
	return b
}
func (b *BIT) reset() {
	b.Tree = make([]int, b.N+1)
}

func (b *BIT) update(idx int) {
	idx = idx + 1
	for idx <= b.N {
		b.Tree[idx]++
		idx += lowBit(idx)
	}
}

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

// 这样解不符和题目的意思
func goodTriplets2(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	bit1 := NewBIT(n + 3)
	bit2 := NewBIT(n + 3)
	for i := 0; i < n; i++ {
		nums1[i] = nums1[i] + 1
		nums2[i] = nums2[i] + 1
	}
	left := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		ch1, ch2 := nums1[i], nums2[i]
		// 找小于的数字个数
		left[i] = min(bit1.query(ch1-1), bit2.query(ch2-1))
		bit1.update(ch1)
		bit2.update(ch2)
	}
	bit1.reset()
	bit2.reset()
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		ch1, ch2 := nums1[i], nums2[i]
		// 找大于的数字个数
		a := bit1.query(n+1) - bit1.query(ch1)
		b := bit2.query(n+1) - bit2.query(ch2)
		right[i] = min(a, b)
		ans += left[i] * min(a, b)
		bit1.update(ch1)
		bit2.update(ch2)
	}
	return int64(ans)
}
