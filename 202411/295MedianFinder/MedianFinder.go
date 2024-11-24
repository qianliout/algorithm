package main

import (
	"container/heap"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

type MedianFinder struct {
	Mx MaxHeap // 左边
	Mi MinHeap // 右边
	// 左边最多比右边多一个
}

func Constructor() MedianFinder {
	s := MedianFinder{
		Mx: make(MaxHeap, 0),
		Mi: make(MinHeap, 0),
	}
	return s
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.Mx, num)
	// 维护
	for {
		has := false
		if this.Mx.Len()-this.Mi.Len() > 1 {
			pop := heap.Pop(&this.Mx).(int)
			heap.Push(&this.Mi, pop)
			has = true
		}
		if len(this.Mx) > 0 && len(this.Mi) > 0 && this.Mx[0] > this.Mi[0] {
			pop := heap.Pop(&this.Mi).(int)
			heap.Push(&this.Mx, pop)
			has = true
		}
		if !has {
			break
		}
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.Mx.Len() == 0 {
		return 0
	}

	if this.Mx.Len() > this.Mi.Len() {
		return float64(this.Mx[0])
	}

	return float64(this.Mx[0]+this.Mi[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
