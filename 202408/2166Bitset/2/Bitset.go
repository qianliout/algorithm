package main

import (
	"fmt"
)

func main() {
	fmt.Println('0', '0'^1, '0'^1^1)
}

type Bitset struct {
	Data []byte
	Cnt  int
	Size int
	flip byte
}

func Constructor(size int) Bitset {
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = '0'
	}
	return Bitset{
		Data: data,
		Cnt:  0,
		Size: size,
		flip: '0',
	}
}

// 将下标为 idx 的位上的值更新为 1 。如果值已经是 1 ，则不会发生任何改变。
func (this *Bitset) Fix(idx int) {
	if idx < 0 || idx >= len(this.Data) {
		return
	}
	if this.flip == this.Data[idx] {
		this.Data[idx] ^= 1
		this.Cnt++
	}
}

// 将下标为 idx 的位上的值更新为 0 。如果值已经是 0 ，则不会发生任何改变。
func (this *Bitset) Unfix(idx int) {
	if idx < 0 || idx >= len(this.Data) {
		return
	}
	if this.flip != this.Data[idx] {
		this.Data[idx] ^= 1
		this.Cnt--
	}
}

func (this *Bitset) Flip() {
	this.flip ^= 1
	this.Cnt = this.Size - this.Cnt
}

func (this *Bitset) All() bool {
	return this.Cnt == this.Size
}

func (this *Bitset) One() bool {
	return this.Cnt > 0
}

func (this *Bitset) Count() int {
	return this.Cnt
}

func (this *Bitset) ToString() string {
	if this.flip == '1' {
		ret := make([]byte, this.Size)
		for i, ch := range this.Data {
			if ch == '0' {
				ret[i] = '1'
			} else {
				ret[i] = '0'
			}
		}
		return string(ret)
	}
	return string(this.Data)
}
