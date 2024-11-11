package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minimumDifference([]int{7, 9, 5, 8, 1, 3}))
}

func minimumDifference(nums []int) int64 {
	m := len(nums)
	n := m / 3
	pre := make(MaxHeap, 0)
	suf := make(MinHeap, 0)
	preS, sufS := 0, 0
	mx := make([]int, m)
	mi := make([]int, m)
	for i := 0; i < m; i++ {
		l := i
		r := m - 1 - i
		preS += nums[l]
		sufS += nums[r]
		heap.Push(&pre, nums[l])
		heap.Push(&suf, nums[r])
		if i >= n {
			preS -= heap.Pop(&pre).(int)
			sufS -= heap.Pop(&suf).(int)
		}
		mi[l] = preS
		mx[r] = sufS
	}
	ans := 1 << 60
	for i := n - 1; i < 2*n; i++ {
		ans = min(ans, mi[i]-mx[i+1])
	}
	return int64(ans)
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Peek() interface{} {
	if len(*h) > 0 {
		return (*h)[0]
	}
	return 0
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] <= h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
