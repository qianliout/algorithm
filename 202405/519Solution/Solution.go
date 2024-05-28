package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := Constructor(3, 1)
	fmt.Println(s.Flip())
	fmt.Println(s.Flip())
	fmt.Println(s.Flip())
	fmt.Println(s.Flip())
}

type Solution struct {
	Used  map[int]int
	M     int
	N     int
	Total int
}

func Constructor(m int, n int) Solution {

	ss := Solution{
		Used:  map[int]int{},
		M:     m,
		N:     n,
		Total: m * n,
	}
	return ss

}

// 太巧妙了
func (this *Solution) Flip() []int {
	if this.Total <= 0 {
		return []int{}
	}
	ans := make([]int, 0)
	x := rand.Intn(this.Total)
	this.Total--
	if y, ok := this.Used[x]; ok {
		ans = this.Split(y)
	} else {
		ans = this.Split(x)
	}
	if y, ok := this.Used[this.Total]; ok {
		this.Used[x] = y
	} else {
		this.Used[x] = this.Total
	}
	return ans
}

func (this *Solution) Reset() {
	this.Used = make(map[int]int)
	this.Total = this.M * this.N
}

func (this *Solution) Split(n int) []int {
	col := n / this.N
	row := n % this.N
	return []int{col, row}
}
