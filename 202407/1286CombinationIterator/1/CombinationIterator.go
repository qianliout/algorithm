package main

import (
	"fmt"
	"math/bits"
)

func main() {
	c := Constructor("abc", 2)
	fmt.Println(c.Next())
	fmt.Println(c.Next())
	fmt.Println(c.Next())
	fmt.Println(c.Next())
	fmt.Println(c.Next())
}

type CombinationIterator struct {
	Char []byte
	L    int
	Cur  int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	ans := []byte(characters)
	le, ri := 0, len(ans)-1
	for le < ri {
		ans[le], ans[ri] = ans[ri], ans[le]
		le++
		ri--
	}

	c := CombinationIterator{
		Char: ans,
		L:    combinationLength,
	}

	c.Cur = 1<<len(characters) - 1
	return c

}

func (this *CombinationIterator) Next() string {
	// 这里 this.Cur>0 也可以
	for this.Cur > 0 && bits.OnesCount(uint(this.Cur)) != this.L {
		this.Cur--
	}

	ans := make([]byte, 0)

	for i := 0; i < len(this.Char); i++ {
		if (this.Cur & (1 << i) >> i) > 0 {
			ans = append(ans, this.Char[i])
		}
	}

	this.Cur--

	le, ri := 0, len(ans)-1
	for le < ri {
		ans[le], ans[ri] = ans[ri], ans[le]
		le++
		ri--
	}

	return string(ans)
}

func (this *CombinationIterator) HasNext() bool {
	// 这里 this.Cur>0就不行，会出错
	for this.Cur >= 0 && bits.OnesCount(uint(this.Cur)) != this.L {
		this.Cur--
	}
	return this.Cur >= 0
}

func countOne(n int) int {
	cnt := 0
	for n != 0 {
		cnt++
		n = (n - 1) & n
	}
	return cnt
}
