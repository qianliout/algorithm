package main

import (
	"container/list"
)

func main() {

}

type LRUCache struct {
	Cap  int
	Head *list.List
	Data map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Cap:  capacity,
		Head: list.New(),
		Data: map[int]*list.Element{},
	}
	return lru
}

type Node struct {
	key   int
	value int
}

func (this *LRUCache) Get(key int) int {
	el, ok := this.Data[key]
	if !ok {
		return -1
	}
	no := el.Value.(*Node)
	// 这样写 el 的地址不会变
	// this.Head.MoveToFront(el)

	this.Head.Remove(el)

	front := this.Head.PushFront(no)
	this.Data[key] = front // 这里是最容易出错的，如是上面是重新插入的，则地址就变了，得重新更新 key 的对应值
	return no.value
}

func (this *LRUCache) Put(key int, value int) {
	el, ok := this.Data[key]
	if ok {
		no := &Node{key: key, value: value}
		front := this.Head.PushFront(no)
		this.Data[key] = front
		this.Head.Remove(el)
		return
	}

	if this.Cap > this.Head.Len() {
		no := &Node{key: key, value: value}
		front := this.Head.PushFront(no)
		this.Data[key] = front
		return
	}

	back := this.Head.Back()
	backN := back.Value.(*Node)
	this.Head.Remove(back)
	delete(this.Data, backN.key)

	no := &Node{key: key, value: value}
	front := this.Head.PushFront(no)
	this.Data[key] = front
}
