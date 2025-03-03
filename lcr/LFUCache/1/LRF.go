package main

import (
	"container/list"
)

func main() {

}

type Node struct {
	key   int
	value int
	freq  int
}

func NewNode(key, value int, freq int) *Node {
	return &Node{
		key:   key,
		value: value,
		freq:  freq,
	}
}

type LFUCache struct {
	Data    map[int]*list.Element
	Freq    map[int]*list.List
	Cap     int
	MinFreq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Data: make(map[int]*list.Element),
		Freq: make(map[int]*list.List),
		Cap:  capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	el, ok := this.Data[key]
	if !ok {
		return -1
	}
	node := el.Value.(*Node)
	if node.freq == this.MinFreq && this.Freq[this.MinFreq].Len() == 1 {
		this.MinFreq++
	}

	this.Freq[node.freq].Remove(el)
	node.freq++
	if this.Freq[node.freq] == nil {
		this.Freq[node.freq] = list.New()
	}
	//  移除时是移除最后面，所以新加的就要加到最前面
	front := this.Freq[node.freq].PushFront(node)
	this.Data[key] = front // key 没有变，直接更新值
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	el, ok := this.Data[key]
	if ok {
		node := el.Value.(*Node)
		if node.freq == this.MinFreq && this.Freq[this.MinFreq].Len() == 1 {
			// 这里直接把 MinFreq++了，会不会导指这个 List 是nil呢，
			// 更新 minFreq
			// 假如MinFreq=1，这个链表里只有一个元素，那么更新后MinFreq 就等于2
			// 会不会存在 Freq=2的链表为空呢？不会，因为 no会马上更新
			this.MinFreq++

		}
		this.Freq[node.freq].Remove(el)
		node.freq++
		if this.Freq[node.freq] == nil {
			this.Freq[node.freq] = list.New()
		}
		node.value = value
		//  移除时是移除最后面，所以新加的就要加到最前面
		front := this.Freq[node.freq].PushFront(node)
		this.Data[key] = front
		return
	}
	if len(this.Data) >= this.Cap {
		ba := this.Freq[this.MinFreq].Back()
		no := ba.Value.(*Node)
		delete(this.Data, no.key)
		// 这一步容易忘记
		this.Freq[this.MinFreq].Remove(ba)
	}

	this.MinFreq = 1
	node := NewNode(key, value, 1)
	if this.Freq[this.MinFreq] == nil {
		this.Freq[this.MinFreq] = list.New()
	}
	front := this.Freq[node.freq].PushFront(node)
	this.Data[key] = front
}
