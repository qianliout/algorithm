package main

import (
	"container/heap"
	"fmt"
	"slices"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1))
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	// 排序不影响原数组
	ids := make([]int, len(nums1))
	for i := range ids {
		ids[i] = i
	}
	// 按nums2从大到小
	slices.SortFunc(ids, func(i, j int) int { return nums2[j] - nums2[i] })
	// 这种方式排序是错的，还不知道原因
	// sort.Slice(ids, func(i, j int) bool { return nums2[i] > nums2[j] })

	mh := make(MinHeap, 0)
	ans, sum := 0, 0

	for _, i := range ids {
		x := nums1[i]
		heap.Push(&mh, x)
		sum += x
		for mh.Len() > k {
			pop := heap.Pop(&mh).(int)
			sum -= pop
		}
		if mh.Len() == k {
			ans = max(ans, nums2[i]*sum)
		}
	}
	return int64(ans)
}
