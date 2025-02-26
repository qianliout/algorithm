package main

import (
	"container/heap"
	. "outback/algorithm/common/commonHeap"
)

func main() {

}

type Checkout struct {
	St     []int
	HeapMx MaxHeap
	Used   map[int]int
}

func Constructor() Checkout {
	ch := Checkout{
		St:     make([]int, 0),
		HeapMx: make(MaxHeap, 0),
		Used:   map[int]int{},
	}
	return ch
}

func (this *Checkout) Get_max() int {
	for this.HeapMx.Len() > 0 {
		pik := this.HeapMx[0]
		if this.Used[pik] > 0 {
			return pik
		}
		heap.Pop(&this.HeapMx)
	}
	return -1
}

func (this *Checkout) Add(value int) {
	this.St = append(this.St, value)
	this.Used[value]++
	heap.Push(&this.HeapMx, value)
}

func (this *Checkout) Remove() int {
	if len(this.St) == 0 {
		return -1
	}
	left := this.St[0]
	this.Used[left]--
	this.St = this.St[1:]
	return left
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */
