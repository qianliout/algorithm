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
	Res []int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{Res: encoding}
}

func (this *RLEIterator) Next(n int) int {
	if n <= 0 || len(this.Res) == 0 {
		return -1
	}

	for {
		if len(this.Res) == 0 {
			return -1
		}
		if n > this.Res[0] {
			n -= this.Res[0]
			this.Res = this.Res[2:]
		} else {
			this.Res[0] -= n
			return this.Res[1]
		}
	}
}
