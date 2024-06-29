package main

import (
	"container/heap"
	"fmt"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	c := Constructor()
	c.AddBack(2)
	fmt.Println(c.PopSmallest())
	fmt.Println(c.PopSmallest())
	fmt.Println(c.PopSmallest())
	c.AddBack(1)
	fmt.Println(c.PopSmallest())
	fmt.Println(c.PopSmallest())
	fmt.Println(c.PopSmallest())
}

type SmallestInfiniteSet struct {
	Mi   MinHeap      // AddBack 回来的元素
	Used map[int]bool // AddBack的元素会有重复
	Idx  int          // [0:idx)是已经 pop 出去的, [idx:] 是还没有 pop 的
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
	// this.Mi 里有元素，说明有元素 AddBack了
	if this.Mi.Len() > 0 {
		pop := heap.Pop(&this.Mi).(int)
		this.Used[pop] = false
		ans = pop
	} else {
		ans = this.Idx
		this.Idx++
	}
	return ans
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	// 增加的元素在原来的无限集中
	if num >= this.Idx {
		return
	}
	// 重复增加
	if this.Used[num] {
		return
	}
	heap.Push(&this.Mi, num)
	this.Used[num] = true
}
