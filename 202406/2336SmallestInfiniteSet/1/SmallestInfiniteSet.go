package main

import (
	"container/heap"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

/*
使用 idx 代表顺序弹出的集合左边界，[idx,+∞] 范围内的数均为待弹出，起始有 idx=1。

考虑当调用 addBack 往集合添加数值 x 时，该如何处理：

	x≥idx：说明数值本身就存在于集合中，忽略该添加操作；
	x=idx−1：数值刚好位于边界左侧，更新 idx=idx−1；
	x<idx−1：考虑将数值添加到某个容器中，该容器支持返回最小值，容易联想到“小根堆”；但小根堆并没有“去重”功能，为防止重复弹出，还需额外使用“哈希表”来记录哪些元素在堆中。

该做法本质上将集合分成两类，一类是从 idx 到正无穷的连续段，对此类操作的复杂度为 O(1)；一类是比 idx 要小的离散类数集，对该类操作复杂度为 O(logn)，其中 n 为调用 addBack 的最大次数。
*/
type SmallestInfiniteSet struct {
	Mi   MinHeap      // AddBack 回来的元素
	Used map[int]bool // AddBack的元素会有重复
	Idx  int          //  [0:idx)是已经 pop 出去的, [idx:] 是还没有 pop 的
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		Mi:   make(MinHeap, 0),
		Used: make(map[int]bool),
		Idx:  1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	ans := -1
	if this.Mi.Len() > 0 {
		pop := heap.Pop(&this.Mi).(int)
		this.Used[pop] = false
		ans = pop
	} else {
		this.Idx++
		ans = this.Idx
	}
	return ans
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.Idx {
		return
	}
	if this.Used[num] {
		return
	}
	// 这个逻辑可以简化
	if num == this.Idx-1 {
		this.Idx--
		return
	}
	heap.Push(&this.Mi, num)
	this.Used[num] = true
}
