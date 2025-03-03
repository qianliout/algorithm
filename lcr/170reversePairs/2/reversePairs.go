package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reversePairs([]int{7, 5, 9, 4, 2}))
	fmt.Println(reversePairs2([]int{7, 5, 9, 4, 2}))
}

// 暴力解法,时间复杂度高
func reversePairs2(record []int) int {
	cnt := 0
	n := len(record)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if record[i] > record[j] {
				cnt++
			}
		}
	}
	return cnt
}

// 树状树组
func reversePairs(record []int) int {
	ans := 0
	n := len(record)
	tem := append([]int{}, record...)
	sort.Ints(tem)
	for i := range record {
		record[i] = sort.SearchInts(tem, record[i])
	}
	bit := NewBIT(n)
	for i := n - 1; i >= 0; i-- {
		ans += bit.Query(record[i] - 1)
		bit.Update(record[i])
	}

	return ans
}

// 「树状数组」是一种可以动态维护序列前缀和的数据结构，它的功能是：
// 单点更新update(i, v)：把序列i位置的数加上一个值v，这题v=1
// 区间查询query(i)：查询序列[1⋯i]区间的区间和，即i位置的前缀和
// 修改和查询的时间代价都是O(logn)，其中n为需要维护前缀和的序列的长度。
// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/solutions/216984/shu-zu-zhong-de-ni-xu-dui-by-leetcode-solution/?envType=problem-list-v2&envId=8LSpuXqD

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

// 值域更新
// 这里是把Data[idx] 里的增加1,那么 Data[idx+1:]的数字都得加1
// 因为是计算前缀和，所以树状树组的idx要加1
func (b *BIT) Update(idx int) {
	idx = idx + 1

	for idx <= b.N {
		b.Tree[idx]++
		idx += lowBit(idx)
	}
}

func (b *BIT) Query(idx int) int {
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
