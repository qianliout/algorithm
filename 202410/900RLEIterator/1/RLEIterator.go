package main

import (
	"fmt"
)

func main() {
	rle := Constructor([]int{3, 8, 0, 9, 2, 5})
	fmt.Println(rle.Next(2))
	fmt.Println(rle.Next(1))
	fmt.Println(rle.Next(1))
	fmt.Println(rle.Next(2))
}

type RLEIterator struct {
	Res []pair
}

func Constructor(encoding []int) RLEIterator {
	res := make([]pair, 0)
	for i := 0; i+1 < len(encoding); i += 2 {
		if encoding[i] == 0 {
			continue
		}
		res = append(res, pair{encoding[i+1], encoding[i]})
	}
	return RLEIterator{Res: res}
}

func (this *RLEIterator) Next(n int) int {
	if n <= 0 {
		return -1
	}
	n = n - 1

	for n > 0 && len(this.Res) > 0 {
		if this.Res[0].Repeat <= n {
			n = n - this.Res[0].Repeat
			this.Res = this.Res[1:]
		} else {
			this.Res[0].Repeat -= n
			n = 0
		}
	}
	if len(this.Res) > 0 {
		ans := this.Res[0].Value
		this.Res[0].Repeat--
		if this.Res[0].Repeat == 0 {
			this.Res = this.Res[1:]
		}
		return ans
	}
	return -1
}

type pair struct {
	Value  int
	Repeat int
}
