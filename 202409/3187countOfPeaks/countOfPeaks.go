package main

import "fmt"

func main() {
	fmt.Println(countOfPeaks([]int{3, 1, 4, 2, 5}, [][]int{{2, 3, 4}, {1, 0, 4}}))
	fmt.Println(countOfPeaks([]int{4, 10, 8, 6}, [][]int{{1, 0, 3}, {1, 2, 3}, {1, 2, 3}}))
}

type BIT struct {
	N    int // 数组长度
	Data []int
}

func NewBIT(n int) *BIT {
	return &BIT{N: n, Data: make([]int, n)}
}

// 更新树状数组中的某个元素
func (b *BIT) update(i, x int) {
	for i < b.N {
		b.Data[i] += x
		i += i & -i
	}
}

func (b *BIT) pre(i int) int {
	res := 0
	for i > 0 {
		res += b.Data[i]
		i = i & (i - 1)
	}
	return res
}

func (b *BIT) query(l, r int) int {
	if l > r {
		return 0
	}
	return b.pre(r) - b.pre(l-1)
}

func countOfPeaks(nums []int, queries [][]int) []int {
	n := len(nums)
	bit := NewBIT(n)
	update := func(i, x int) {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			bit.update(i, x)
		}
	}
	for i := 1; i < n-1; i++ {
		update(i, 1)
	}

	ans := make([]int, 0)

	for _, q := range queries {
		op, l, r := q[0], q[1], q[2]
		if op == 1 {
			ans = append(ans, bit.query(l+1, r-1))
			continue
		}
		// 先撤销
		for j := max(l-1, 1); j < min(l+2, n-1); j++ {
			update(j, -1)
		}
		nums[l] = r
		for j := max(l-1, 1); j < min(l+2, n-1); j++ {
			update(j, 1)
		}
	}
	return ans
}

func lowBit(x int) int {
	return x & -x
}