package main

import (
	"container/list"
)

func main() {

}

type LRUCache struct {
	Capacity int
	Head     *list.List
	Data     map[int]*list.Element
}
type Node struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Capacity: capacity,
		Head:     list.New(),
		Data:     make(map[int]*list.Element),
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.Data[key]
	if !ok {
		return -1
	}
	// this.Head.Remove(e)
	// back := this.Head.PushBack(e.Value)
	// this.Data[key] = back
	this.Head.MoveToBack(e)
	node := e.Value.(*Node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.Data[key]
	if ok {
		// 去除原来的
		this.Head.Remove(e)
		// 加上新的
		back := this.Head.PushBack(&Node{Key: key, Value: value})
		// 更新 data
		this.Data[key] = back // data中的地址有变动
		return
	}
	// 如果没有,且容量还够
	if len(this.Data) < this.Capacity {
		node := &Node{Key: key, Value: value}
		back := this.Head.PushBack(node)
		this.Data[key] = back

		return
	}
	// 容量也不够了
	front := this.Head.Front()
	fir := this.Head.Remove(front).(*Node)

	delete(this.Data, fir.Key) // 上面是通过Len来判断容量是否够
	node := &Node{Key: key, Value: value}
	back := this.Head.PushBack(node)
	this.Data[key] = back
}

//  最近最少使用
// 实现 LRUCache 类：
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
// 当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
