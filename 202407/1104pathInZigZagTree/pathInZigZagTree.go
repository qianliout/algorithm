package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(pathInZigZagTree(14))
}

func pathInZigZagTree(label int) []int {
	/*
		发现没有，除了第一位都是1，根节点和子节点的所有位相反！
		好了，这就是获取根节点的方法了：去除最后一位，并将除了最高位以外的所有位取反
		我们只需要重复这个过程，就可以获取路径了
	*/
	row := bits.Len(uint(label))
	ans := make([]int, row)
	for i := row - 1; i >= 0; i-- {
		ans[i] = label
		label = (((1 << i) - 1) ^ label) >> 1
	}

	return ans
}
