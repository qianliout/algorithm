package main

import (
	"math"
)

func main() {
	set := Constructor()
	set.Add(1000000)
	set.Contains(1000000)
	set.Contains(1000000000)
}

type MyHashSet struct {
	Bucket  [][]int
	B       int // 每个 bucket可以装多少数据
	Cnt     int // Bucket 的数量
	Deleted int
}

func Constructor() MyHashSet {
	single := 100
	all := int(math.Pow(10, 6))
	data := make([][]int, all/single+1)

	set := MyHashSet{
		Bucket:  data,
		Cnt:     all / single,
		B:       single,
		Deleted: -1,
	}

	return set
}

func (this *MyHashSet) Add(key int) {
	idx := key / this.Cnt
	if len(this.Bucket[idx]) == 0 {
		this.Bucket[idx] = make([]int, this.B)
	}
	idx2 := key % this.B

	this.Bucket[idx][idx2] = key
}

func (this *MyHashSet) Remove(key int) {
	idx := key / this.Cnt
	if len(this.Bucket[idx]) == 0 {
		return
	}
	idx2 := key % this.B
	this.Bucket[idx][idx2] = this.Deleted
}

func (this *MyHashSet) Contains(key int) bool {
	idx := key / this.Cnt
	id2 := key % this.B
	if len(this.Bucket[idx]) == 0 {
		return false
	}
	return this.Bucket[idx][id2] == key
}
