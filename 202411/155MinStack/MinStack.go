package main

import (
	"container/heap"
	. "outback/algorithm/common/commonHeap"
)

func main() {

}

type MinStack struct {
	Hp   MinHeap
	Del  map[int]int
	Data []int
}

func Constructor() MinStack {
	return MinStack{
		Hp:   make(MinHeap, 0),
		Del:  make(map[int]int),
		Data: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.Data = append(this.Data, val)
	heap.Push(&this.Hp, val)
	this.Del[val]++
}

func (this *MinStack) Pop() {
	n := len(this.Data)
	va := this.Data[n-1]
	this.Data = this.Data[:n-1]
	this.Del[va]--
}

func (this *MinStack) Top() int {
	n := len(this.Data)
	return this.Data[n-1]
}

func (this *MinStack) GetMin() int {
	for {
		va := this.Hp[0]
		if this.Del[va] <= 0 {
			heap.Pop(&this.Hp)
		} else {
			return va
		}
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
