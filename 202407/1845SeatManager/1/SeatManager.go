package main

import (
	"container/heap"
	"fmt"
)

func main() {
	c := Constructor(10)
	fmt.Println(c.Reserve())
	fmt.Println(c.Reserve())
	c.Unreserve(2)
	fmt.Println(c.Reserve())
	fmt.Println(c.Reserve())
	fmt.Println(c.Reserve())
	fmt.Println(c.Reserve())
	c.Unreserve(5)
}

type SeatManager struct {
	N    int
	HP   MinHeap
	Seat map[int]*Seat
}

func Constructor(n int) SeatManager {
	s := SeatManager{
		N:    n,
		HP:   make(MinHeap, 0),
		Seat: make(map[int]*Seat),
	}
	return s

}

func (this *SeatManager) Reserve() int {
	// 每一次对 reserve 的调用，题目保证至少存在一个可以预约的座位。
	if len(this.HP) > 0 && this.HP[0].Used == 0 {
		ans := this.HP[0].Idx
		this.HP[0].Used = 1
		heap.Fix(&this.HP, 0)
		return ans
	}
	ans := len(this.HP) + 1
	set := &Seat{
		Idx:  ans,
		Used: 1,
	}
	this.Seat[ans] = set
	heap.Push(&this.HP, set)

	return ans
}

func (this *SeatManager) Unreserve(seatNumber int) {
	// 每一次对 unreserve 的调用，题目保证 seatNumber 在调用函数前都是被预约状态
	set := this.Seat[seatNumber]
	set.Used = 0
	heap.Init(&this.HP)
}

type Seat struct {
	Idx  int
	Used int // 没有使用就是0，使用了就是1
}

type MinHeap []*Seat

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].Used < h[j].Used {
		return true
	} else if h[i].Used > h[j].Used {
		return false
	}
	return h[i].Idx < h[j].Idx
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Seat))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
