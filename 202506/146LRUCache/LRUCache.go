package main

import (
	"container/list"
)

func main() {

}

type LRUCache struct {
	Data map[int]*list.Element
	List *list.List
	Cap  int
}

type Node struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Data: make(map[int]*list.Element),
		List: list.New(),
		Cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	el, ok := this.Data[key]
	if !ok {
		return -1
	}
	va := el.Value.(*Node)
	this.List.MoveToFront(el)
	return va.Value
}

func (this *LRUCache) Put(key int, value int) {
	el, ok := this.Data[key]
	if ok {
		el.Value = &Node{
			Key:   key,
			Value: value,
		}
		this.List.MoveToFront(el)
		return
	}
	if this.List.Len() < this.Cap {
		no := &Node{
			Key:   key,
			Value: value,
		}
		front := this.List.PushFront(no)
		this.Data[key] = front
		return
	}
	back := this.List.Back()
	k := back.Value.(*Node)
	delete(this.Data, k.Key)
	this.List.Remove(back)

	no := &Node{
		Key:   key,
		Value: value,
	}
	front := this.List.PushFront(no)
	this.Data[key] = front
	return
}
