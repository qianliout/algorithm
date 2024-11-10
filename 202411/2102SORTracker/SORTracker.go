package main

import (
	"container/heap"
)

func main() {

}

type SORTracker struct {
	MxHeap MaxHeap
	MiHeap MinHeap
}

func Constructor() SORTracker {
	return SORTracker{
		MxHeap: make(MaxHeap, 0), // 坏班
		MiHeap: make(MinHeap, 0), // 好班
	}
}

func (this *SORTracker) Add(name string, score int) {
	// 要想实现招收学生，但不改变好班人数，那肯定要把新生送到坏班。
	// 但是有个问题，这个新生万一实力很强呢？那岂不是得升到好班，再把好班淘汰一个到坏班
	// 所以我们不管三七二十一，先把新生送到好班去，再淘汰一个送到坏班，这样最方便了
	p := pair{Name: name, Score: score}
	heap.Push(&this.MiHeap, p)
	pop := heap.Pop(&this.MiHeap).(pair)
	heap.Push(&this.MxHeap, pop)
}

func (this *SORTracker) Get() string {
	pop := heap.Pop(&this.MxHeap).(pair)
	heap.Push(&this.MiHeap, pop)
	return pop.Name
}

type pair struct {
	Name  string
	Score int
}

type MaxHeap []pair

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].Score != h[j].Score {
		return h[i].Score > h[j].Score
	}
	return h[i].Name < h[j].Name

}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
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

type MinHeap []pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].Score != h[j].Score {
		return h[i].Score < h[j].Score
	}
	return h[i].Name > h[j].Name // 这里容易出错

}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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

func (h *MinHeap) Peek() interface{} {
	if len(*h) > 0 {
		return (*h)[0]
	}
	return 0
}
