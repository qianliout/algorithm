package main

import (
	"fmt"
	"math/bits"
	"strings"
)

func main() {
	fmt.Println(uint64(1 << 63))
}

// 错误解法
type Bitset struct {
	Data []uint64
	L    int
	N    int // 每个数字存储多少位默认63位
	Size int
}

func Constructor(size int) Bitset {
	b := Bitset{N: 64}
	l := (size + b.N - 1) / b.N
	data := make([]uint64, (size+b.N)/b.N) // 向上取整
	b.Data = data
	b.L = l
	b.Size = size
	return b
}

// 将下标为 idx 的位上的值更新为 1 。如果值已经是 1 ，则不会发生任何改变。
func (this *Bitset) Fix(idx int) {
	n := idx / this.N
	p := idx % this.N
	a := this.Data[n]

	a = a | 1<<p
	this.Data[n] = a

}

// 将下标为 idx 的位上的值更新为 0 。如果值已经是 0 ，则不会发生任何改变。
func (this *Bitset) Unfix(idx int) {
	n := idx / this.N
	p := idx % this.N
	a := this.Data[n]
	a = a ^ (1 << p)
	this.Data[n] = a
}

func (this *Bitset) Flip() {
	for i, ch := range this.Data {
		this.Data[i] = ^ch
	}
}

func (this *Bitset) All() bool {
	for _, ch := range this.Data {
		if bits.OnesCount64(ch) != this.N {
			return false
		}
	}
	return true
}

func (this *Bitset) One() bool {
	for _, ch := range this.Data {
		if ch > 0 {
			return true
		}
	}
	return false
}

func (this *Bitset) Count() int {
	ans := 0
	for _, ch := range this.Data {
		ans += bits.OnesCount64(ch)
	}
	return ans
}

func (this *Bitset) ToString() string {
	ans := []string{}
	for _, ch := range this.Data {
		ans = append(ans, this.toString(ch))
	}

	a := strings.Join(ans, "")
	mi := min(this.Size, len(a))
	return a[:mi]
}

func (this *Bitset) toString(ch uint64) string {
	ans := make([]string, 0)
	i := 0
	for i < this.N {
		ans = append(ans, fmt.Sprintf("%d", ch&1))
		ch = ch >> 1
		i++
	}
	a := strings.Join(ans, "")
	return a

}

func SetBit0(value uint64, flag uint64) uint64 {
	pre := value
	pre &= ^(1 << flag)
	return pre
}

// value从右到左的第flag位设置成1
func SetBit1(value uint64, flag uint64) uint64 {
	pre := value
	pre |= 1 << (flag)
	return pre
}
