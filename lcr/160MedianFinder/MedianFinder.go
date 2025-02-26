package main

import (
	"container/heap"
	"fmt"
	. "outback/algorithm/common/commonHeap"
)

func main() {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	m.AddNum(3)
	fmt.Println(m.FindMedian())
}

type MedianFinder struct {
	// 保证左边大堆和右边小堆的元素个数之差不超过1,且总是左边多
	LeftMx  MaxHeap // 左边大堆
	RightMi MinHeap // 右边小堆
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := MedianFinder{
		LeftMx:  make(MaxHeap, 0),
		RightMi: make(MinHeap, 0),
	}

	return m
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.LeftMx, num)
	for {
		ready := true

		if this.LeftMx.Len()-this.RightMi.Len() > 1 {
			a := heap.Pop(&this.LeftMx).(int)
			heap.Push(&this.RightMi, a)
			ready = false
		}
		if this.RightMi.Len() > this.LeftMx.Len() {
			a := heap.Pop(&this.RightMi).(int)
			heap.Push(&this.LeftMx, a)
			ready = false
		}

		if this.RightMi.Len() > 0 && this.LeftMx.Len() > 0 && this.LeftMx[0] > this.RightMi[0] {
			a := heap.Pop(&this.LeftMx).(int)
			heap.Push(&this.RightMi, a)
			ready = false
		}

		if ready {
			break
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.LeftMx.Len() == 0 && this.RightMi.Len() == 0 {
		return 0
	}
	if this.LeftMx.Len() > this.RightMi.Len() {
		return float64(this.LeftMx[0])
	}
	return float64(this.LeftMx[0]+this.RightMi[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
