package main

import (
	"math"
)

func main() {

}

type Fancy struct {
	Data []int
	Mod  int
}

func Constructor() Fancy {
	return Fancy{
		Data: make([]int, 0),
		Mod:  int(math.Pow10(9)) + 7,
	}
}

func (this *Fancy) Append(val int) {
	this.Data = append(this.Data, val)
}

func (this *Fancy) AddAll(inc int) {
	for i := range this.Data {
		this.Data[i] += inc
		this.Data[i] %= this.Mod
	}
}

func (this *Fancy) MultAll(m int) {
	for i := range this.Data {
		this.Data[i] *= m
		this.Data[i] %= this.Mod
	}
}

func (this *Fancy) GetIndex(idx int) int {
	if idx < 0 || idx >= len(this.Data) {
		return -1
	}
	return this.Data[idx]
}
