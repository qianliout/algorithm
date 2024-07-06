package main

import (
	"fmt"
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
	Char string
	L    int
	Cur  int
}

// 不知道是那里不对

func Constructor(characters string, combinationLength int) CombinationIterator {
	c := CombinationIterator{
		Char: characters,
		L:    combinationLength,
	}

	c.Cur = 1<<len(characters) - 1
	return c

}

func (this *CombinationIterator) Next() string {
	for this.Cur >= 0 && countOne(this.Cur) != this.L {
		this.Cur--
	}
	// if this.Cur <= 0 {
	// 	return ""
	// }

	ans := make([]byte, 0)

	for i := 0; i < len(this.Char); i++ {
		if (this.Cur & (1 << i) >> i) > 0 {
			ans = append(ans, this.Char[i])
		}
	}

	this.Cur--
	// le, ri := 0, len(ans)-1
	// for le < ri {
	// 	ans[le], ans[ri] = ans[ri], ans[le]
	// 	le++
	// 	ri--
	// }

	return string(ans)
}

func (this *CombinationIterator) HasNext() bool {
	for this.Cur > 0 {
		if countOne(this.Cur) != this.L {

			return true
		}
		this.Cur--
	}
	return false
}

func countOne(n int) int {
	cnt := 0
	for n != 0 {
		cnt++
		n = (n - 1) & n
	}
	return cnt
}
