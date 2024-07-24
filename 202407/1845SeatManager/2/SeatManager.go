package main

import (
	"container/heap"
	"fmt"

	. "outback/algorithm/common/commonHeap"
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
	N     int
	HP    MinHeap
	seats int
}

func Constructor(n int) SeatManager {
	s := SeatManager{
		N:  n,
		HP: make(MinHeap, 0),
	}
	return s

}

func (this *SeatManager) Reserve() int {
	if len(this.HP) > 0 {
		pop := heap.Pop(&this.HP).(int)
		return pop
	}
	this.seats += 1 // 添加一把新的椅子
	return this.seats
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&this.HP, seatNumber)
}
