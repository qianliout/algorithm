package main

import (
	"math/rand"
	"time"
)

func main() {

}

type RandomizedSet struct {
	Data map[int]int
	Nums []int
	Rand *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Data: make(map[int]int),
		Nums: make([]int, 0),
		Rand: rand.New(rand.NewSource(time.Now().UnixMilli())),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Data[val]; !ok {
		this.Nums = append(this.Nums, val)
		this.Data[val] = len(this.Nums) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.Data[val]; ok {
		idx := this.Data[val]
		last := len(this.Nums) - 1
		this.Nums[idx] = this.Nums[last]
		this.Data[this.Nums[idx]] = idx
		this.Nums = this.Nums[:len(this.Nums)-1]
		delete(this.Data, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	n := this.Rand.Intn(len(this.Nums))
	return this.Nums[n]
}
