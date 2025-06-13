package main

import (
	"container/heap"

	. "outback/algorithm/common/commonHeap"
	. "outback/algorithm/common/listnode"
)

func main() {

}

type MedianFinder struct {
	left  MaxHeap // left
	right MinHeap // right
}

func Constructor() MedianFinder {
	return MedianFinder{
		right: make(MinHeap, 0),
		left:  make(MaxHeap, 0),
	}
}

// 原则，左边最多比右边多一个
// 先加到左边
func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.left, num)

	for {
		find := false
		if this.left.Len()-this.right.Len() > 1 {
			pop := heap.Pop(&this.left)
			heap.Push(&this.right, pop)
			find = true
		}
		if this.right.Len() > 0 && this.left.Len() > 0 && this.left[0] > this.right[0] {
			pop := heap.Pop(&this.left)
			heap.Push(&this.right, pop)
			find = true
		}

		if this.right.Len()-this.left.Len() > 0 {
			pop := heap.Pop(&this.right)
			heap.Push(&this.left, pop)
			find = true
		}

		if !find {
			break
		}
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return float64(this.left[0])
	} else if this.left.Len() == 0 {
		return 0
	}
	return float64(this.left[0]+this.right[0]) / float64(2)
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 1 {
		return lists[0]
	}
	if n == 0 {
		return nil
	}
	mid := n / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == l1 {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
	l2.Next = merge(l1, l2.Next)
	return l2
}
