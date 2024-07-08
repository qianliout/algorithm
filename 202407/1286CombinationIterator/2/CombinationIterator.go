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
	ChanLength int
	Char       []byte
	L          int
	Cur        int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	ans := []byte(characters)
	c := CombinationIterator{
		ChanLength: len(ans),
		Char:       ans,
		L:          combinationLength,
	}

	c.Cur = 1<<len(characters) - 1
	return c

}

func (this *CombinationIterator) Next() string {
	// 这里 this.Cur>0 也可以
	for this.Cur > 0 && bits.OnesCount(uint(this.Cur)) != this.L {
		this.Cur--
	}
	if this.Cur <= 0 {
		// 说明没有值了，但是题目中说测试用例中不会有这样的例子
		return ""
	}
	ans := make([]byte, 0)

	for i := 0; i < len(this.Char); i++ {
		// char中的第i位，bit 是反的,假如：abc，对应的 二制是111,a在abc 中是第0位，但是在进进制中是从左到右的第二位
		left := this.ChanLength - 1 - i
		if (this.Cur & (1 << left) >> left) > 0 {
			ans = append(ans, this.Char[i])
		}
	}

	this.Cur--
	return string(ans)
}

func (this *CombinationIterator) HasNext() bool {
	for this.Cur > 0 && bits.OnesCount(uint(this.Cur)) != this.L {
		this.Cur--
	}
	return this.Cur > 0
}
