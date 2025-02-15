package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(getNumberOfBacklogOrders([][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}))
}

func getNumberOfBacklogOrders(orders [][]int) int {
	sellHeap := make(MinHeap, 0)
	buyHeap := make(MaxHeap, 0)
	for i, ch := range orders {
		// orders[i] = [pricei, amounti, orderTypei]
		or := Order{
			price:     ch[0],
			amount:    ch[1],
			orderType: ch[2],
			idx:       i,
		}
		// sell
		if or.orderType == 1 {
			for buyHeap.Len() > 0 && or.amount > 0 {
				if or.price > buyHeap[0].price {
					break
				}
				pop := heap.Pop(&buyHeap).(Order)
				if pop.amount == or.amount {
					or.amount -= pop.amount
				} else if pop.amount > or.amount {
					pop.amount = pop.amount - or.amount
					or.amount = 0
					heap.Push(&buyHeap, pop)
				} else if pop.amount < or.amount {
					or.amount -= pop.amount
				}
			}
			if or.amount > 0 {
				heap.Push(&sellHeap, or)
			}
		}
		// buy
		if or.orderType == 0 {
			for sellHeap.Len() > 0 && or.amount > 0 {
				if sellHeap[0].price > or.price {
					break
				}
				pop := heap.Pop(&sellHeap).(Order)
				if pop.amount == or.amount {
					or.amount -= pop.amount
				} else if pop.amount > or.amount {
					pop.amount = pop.amount - or.amount
					or.amount = 0
					heap.Push(&sellHeap, pop)
				} else if pop.amount < or.amount {
					or.amount -= pop.amount
				}
			}
			if or.amount > 0 {
				heap.Push(&buyHeap, or)
			}
		}
	}
	ans := 0
	mod := 1000000007
	for i := range buyHeap {
		ans += buyHeap[i].amount
		ans = ans % mod
	}
	for i := range sellHeap {
		ans += sellHeap[i].amount
		ans = ans % mod
	}
	return ans
}

type Order struct {
	price     int
	amount    int
	orderType int
	idx       int
}

type MaxHeap []Order

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price > h[j].price
	}
	return h[i].idx < h[j].idx
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Order))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []Order

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price < h[j].price
	}
	return h[i].idx < h[j].idx
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Order))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
