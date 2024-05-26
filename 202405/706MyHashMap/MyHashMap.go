package main

import (
	"math"
)

func main() {

}

func (this *MyHashMap) Put(key int, value int) {
	this.Data[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if this.Data[key] == this.Delete || this.Data[key] == this.NotAdd {
		return -1
	}
	return this.Data[key]
}

type MyHashMap struct {
	Data   []int
	Delete int
	NotAdd int
}

func Constructor() MyHashMap {
	cnt := int(math.Pow(10, 6)) + 1

	s := MyHashMap{
		Data:   make([]int, cnt),
		Delete: -2,
		NotAdd: -1,
	}
	for i := range s.Data {
		s.Data[i] = s.NotAdd
	}
	return s
}

func (this *MyHashMap) Remove(key int) {
	this.Data[key] = this.Delete
}
