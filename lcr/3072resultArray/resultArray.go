package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(resultArray([]int{2, 1, 3, 3}))
	fmt.Println(resultArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(resultArray([]int{5, 14, 3, 1, 2})) // 5,3,1,2,14
}

func resultArray(nums []int) []int {
	n := len(nums)
	tem := append([]int{}, nums...)
	sort.Ints(tem)
	record := make([]int, n)
	// 离散化
	for i := range nums {
		record[i] = sort.SearchInts(tem, nums[i])
	}
	ans1 := make([]int, 0)
	ans2 := make([]int, 0)
	bit1 := NewBit(n + 2)
	bit2 := NewBit(n + 2)
	ans1 = append(ans1, nums[0])
	ans2 = append(ans2, nums[1])

	bit1.update(sort.SearchInts(tem, nums[0]))
	bit2.update(sort.SearchInts(tem, nums[1]))

	for i := 2; i < len(nums); i++ {
		ch := sort.SearchInts(tem, nums[i])
		a := bit1.query(n) - bit1.query(ch)
		b := bit2.query(n) - bit2.query(ch)
		if a > b {
			ans1 = append(ans1, nums[i])
			bit1.update(ch)
		} else if a < b {
			ans2 = append(ans2, nums[i])
			bit2.update(ch)
		} else {
			if len(ans1) <= len(ans2) {
				ans1 = append(ans1, nums[i])
				bit1.update(ch)
			} else {
				ans2 = append(ans2, nums[i])
				bit2.update(ch)
			}
		}
	}
	ans1 = append(ans1, ans2...)
	return ans1
}

type BIT struct {
	N    int
	Tree []int
}

func NewBit(n int) *BIT {
	b := &BIT{
		N:    n,
		Tree: make([]int, n+1),
	}
	return b
}

// 计算大于登录 idx 的值有多少个
func (b *BIT) query(idx int) int {
	idx = idx + 1
	ans := 0
	for idx > 0 {
		ans += b.Tree[idx]
		idx -= lowBit(idx)
	}

	return ans
}

// idx 处理的值增加1，那么后面的所有值都要增加1
func (b *BIT) update(idx int) {
	idx = idx + 1
	for idx <= b.N {
		b.Tree[idx]++
		idx += lowBit(idx)
	}
}

func lowBit(n int) int {
	return n & (-n)
}
