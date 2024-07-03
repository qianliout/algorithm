package main

import (
	"container/heap"
	"fmt"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	dp := Constructor(2)
	dp.Push(1)
	dp.Push(2)
	dp.Push(3)
	dp.Push(4)
	dp.Push(5)
	fmt.Println(dp.PopAtStack(0))
	dp.Push(20)
	dp.Push(21)
	fmt.Println(dp.PopAtStack(0))
	fmt.Println(dp.PopAtStack(2))
	fmt.Println(dp.Pop())
	fmt.Println(dp.Pop())
	fmt.Println(dp.Pop())
	fmt.Println(dp.Pop())
	fmt.Println(dp.Pop())

	a := make([][]int, 1)
	a[0] = []int{1, 2, 3}
	st := a[0]
	st = st[:len(st)-1]
	fmt.Println(a)
	a[0] = st
	fmt.Println(a)

}

type DinnerPlates struct {
	Capacity int
	Stark    [][]int
	MH       MinHeap
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		Capacity: capacity,
		Stark:    make([][]int, 0),
		MH:       make(MinHeap, 0),
	}

}

func (this *DinnerPlates) Push(val int) {
	if this.MH.Len() > 0 && this.MH[0] >= len(this.Stark) {
		this.MH = make(MinHeap, 0)
	}

	pushed := false
	for this.MH.Len() > 0 {
		index := heap.Pop(&this.MH).(int)
		if index >= len(this.Stark) || index < 0 {
			continue
		}
		if len(this.Stark[index]) == this.Capacity {
			continue
		}
		pushed = true
		this.Stark[index] = append(this.Stark[index], val)
		// st = append(st, val) // 这样写会有两次赋值 会导致超时
		if len(this.Stark[index]) < this.Capacity {
			heap.Push(&this.MH, index)
		}
		return
	}

	if !pushed {
		this.Stark = append(this.Stark, []int{val})
		if this.Capacity > 1 {
			heap.Push(&this.MH, len(this.Stark)-1)
		}
	}
}

func (this *DinnerPlates) Pop() int {
	return this.PopAtStack(len(this.Stark) - 1)
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index >= len(this.Stark) || index < 0 {
		return -1
	}
	if len(this.Stark[index]) == 0 {
		return -1
	}
	n := len(this.Stark[index])
	if n == this.Capacity { // 原来是满的，现在pop 一个之后就不满了
		heap.Push(&this.MH, index)
	}
	va := this.Stark[index][n-1]
	this.Stark[index] = this.Stark[index][:n-1]
	// 这样写会重复 push,会导致超时
	// if len(this.Stark[index]) < this.Capacity {
	// 	heap.Push(&this.MH, index)
	// }
	// 把后面的空栈删除,没有这个优化，会导致上面 push 时会添加到最后的空栈，进而pop 时会检测很多空栈
	for len(this.Stark) > 0 && len(this.Stark[len(this.Stark)-1]) == 0 {
		this.Stark = this.Stark[:len(this.Stark)-1]
	}
	return va
}
