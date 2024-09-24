package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumSumQueries([]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, [][]int{{4, 1}, {1, 3}, {2, 5}}))

}

type BinaryIndexedTree struct {
	n int
	c []int
}

func NewBinaryIndexedTree(n int) BinaryIndexedTree {
	c := make([]int, n+1)
	for i := range c {
		c[i] = -1
	}
	return BinaryIndexedTree{n: n, c: c}
}

func (bit *BinaryIndexedTree) update(x, v int) {
	for x <= bit.n {
		bit.c[x] = max(bit.c[x], v)
		x += x & -x
	}
}

func (bit *BinaryIndexedTree) query(x int) int {
	mx := -1
	for x > 0 {
		mx = max(mx, bit.c[x])
		x -= x & -x
	}
	return mx
}

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	n, m := len(nums1), len(queries)
	nums := make([][2]int, n)
	for i := range nums {
		nums[i] = [2]int{nums1[i], nums2[i]}
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i][0] > nums[j][0] })
	sort.Ints(nums2)
	ids := make([]int, m)
	for i := range ids {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return queries[ids[j]][0] < queries[ids[i]][0]
	})

	tree := NewBinaryIndexedTree(n)
	ans := make([]int, m)
	j := 0
	for _, i := range ids {
		x, y := queries[i][0], queries[i][1]
		for ; j < n && nums[j][0] >= x; j++ {
			k := n - sort.SearchInts(nums2, nums[j][1])
			tree.update(k, nums[j][0]+nums[j][1])
		}
		k := n - sort.SearchInts(nums2, y)
		ans[i] = tree.query(k)
	}
	return ans
}

// 遍历每个查询 queries[i]=(x,y)，对于当前查询，我们循环将 nums 中所有大于等于 x 的元素的 nums2
// 的值插入到树状数组中，树状数组维护的是离散化后的 nums2 的区间中 nums1+nums2 的最大值。那么我们只需要在树状数组中查询大于等于离散化后的 y
// 区间对应的最大值即可。注意，由于树状数组维护的是前缀最大值，所以我们在实现上，可以将 nums2 反序插入到树状数组中。
