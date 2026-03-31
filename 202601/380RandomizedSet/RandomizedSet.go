package main

import (
	"math/rand"
	"time"
)

func main() {}

type RandomizedSet struct {
	Data []int
	Exit map[int]int
	Rand *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Data: make([]int, 0),
		Exit: make(map[int]int),
		Rand: rand.New(rand.NewSource(time.Now().UnixMilli())),
	}

}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.Exit[val]
	if ok {
		return false
	}
	this.Data = append(this.Data, val)
	this.Exit[val] = len(this.Data) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	l, ok := this.Exit[val]
	if !ok {
		return false
	}
	last := len(this.Data) - 1
	this.Exit[this.Data[last]] = l
	this.Data[l], this.Data[last] = this.Data[last], this.Data[l]
	this.Data = this.Data[:last]
	delete(this.Exit, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	n := this.Rand.Int31n(int32(len(this.Data)))
	return this.Data[int(n)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
