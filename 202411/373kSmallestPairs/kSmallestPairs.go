package main

import (
	"container/heap"
)

func main() {

}

// 这样做是错的
// [1,2,4,5,6]
// [3,5,7,9]
// k = 3
func kSmallestPairs1(nums1 []int, nums2 []int, k int) [][]int {
	ans := make([][]int, 0)
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	for i < m && j < n && len(ans) < k {
		ans = append(ans, []int{nums1[i], nums2[j]})
		if nums1[i] < nums2[j] {
			j++
		} else {
			i++
		}
	}
	return ans
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	ans := make([][]int, 0)
	m, n := len(nums1), len(nums2)
	hm := make(MinHeap, 0)
	for i := 0; i < m; i++ {
		heap.Push(&hm, pair{a: i, b: 0, sum: nums1[i] + nums2[0]})
	}
	for len(ans) < k && hm.Len() > 0 {
		pop := heap.Pop(&hm).(pair)
		i, j := pop.a, pop.b
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&hm, pair{a: i, b: j + 1, sum: nums1[i] + nums2[j+1]})
		}
	}
	return ans
}

type pair struct {
	a, b int
	sum  int
}

type MinHeap []pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
