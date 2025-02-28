package main

import (
	"container/list"
)

func main() {

}

type LFUCache struct {
	Capacity int
	Data     map[int]*list.Element
	Freq     map[int]*list.List // 使用频率
	Exit     map[int]int
}

type Node struct {
	Key   int
	Value int
	Freq  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{}
}

func (this *LFUCache) Get(key int) int {
	e, ok := this.Data[key]
	if !ok {
		return -1
	}
	node := e.Value.(*Node)
	le := this.Freq[node.Freq]
	le.Remove(e)
	// if le.Len() == 0 {
	//
	// }
	node.Freq++

	if this.Freq[node.Freq] == nil {
		this.Freq[node.Freq] = list.New()
	}
	this.Freq[node.Freq].PushBack(node)

	return node.Value
}

func (this *LFUCache) Put(key int, value int) {
	e, ok := this.Data[key]
	if ok {
		node := e.Value.(*Node)

	}

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

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
