package main

import (
	"math"
)

func main() {

}

type MyHashSet struct {
	Data   []int
	Delete int
	Added  int
}

func Constructor() MyHashSet {
	cnt := int(math.Pow(10, 6)) + 1
	return MyHashSet{
		Data:   make([]int, cnt),
		Delete: -1,
		Added:  1,
	}

}

func (this *MyHashSet) Add(key int) {
	this.Data[key] = this.Added
}

func (this *MyHashSet) Remove(key int) {
	this.Data[key] = this.Delete
}

func (this *MyHashSet) Contains(key int) bool {
	return this.Data[key] == this.Added
}
