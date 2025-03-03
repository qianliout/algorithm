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
	no := el.Value.(*Node)
	this.Freq[no.freq].Remove(el)
	// 更新 minFreq
	// 假如MinFreq=1，这个链表里只有一个元素，那么更新后MinFreq 就等于2
	// 会不会存在 Freq=2的链表为空呢？不会，因为 no会马上更新
	if this.MinFreq == no.freq && this.Freq[no.freq].Len() == 0 {
		this.MinFreq = no.freq + 1
	}
	no.freq++
	if this.Freq[no.freq] == nil {
		this.Freq[no.freq] = list.New()
	}
	front := this.Freq[no.freq].PushFront(no)
	this.Data[key] = front

	return no.value
}

func (this *LFUCache) Put(key int, value int) {
	if el, ok := this.Data[key]; ok {
		no := el.Value.(*Node)
		this.Freq[no.freq].Remove(el)
		if this.MinFreq == no.freq && this.Freq[no.freq].Len() == 0 {
			this.MinFreq = no.freq + 1
		}

		no.freq++
		no.value = value
		if this.Freq[no.freq] == nil {
			this.Freq[no.freq] = list.New()
		}
		front := this.Freq[no.freq].PushFront(no)
		this.Data[key] = front
		return
	}

	if len(this.Data) >= this.Cap {
		back := this.Freq[this.MinFreq].Back()
		backN := back.Value.(*Node)
		this.Freq[this.MinFreq].Remove(back)
		delete(this.Data, backN.key)
	}
	this.MinFreq = 1
	no := NewNode(key, value, 1)
	if this.Freq[no.freq] == nil {
		this.Freq[no.freq] = list.New()
	}
	front := this.Freq[no.freq].PushFront(no)
	this.Data[key] = front
	return
}

/*
则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最久未使用 的键。
为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
*/
