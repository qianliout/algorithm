package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(minimumLines([][]int{{1, 3}, {2, 3}, {3, 3}}))
}

func minimumLines(stockPrices [][]int) int {
	sort.Slice(stockPrices, func(i, j int) bool { return stockPrices[i][0] < stockPrices[j][0] })
	preDY, preDx := 1, 0 // 这是难点，但是为啥要这样写呢，这样可以认为第一个点的斜率是无穷大
	// preDY, preDx := 0, 1 //  这样写就不行，因为这样写，可以认为第一个点的斜率是0，斜率是0，是存在的，如果有其他点的斜率是0，就会出错
	ans := 0
	for i := 1; i < len(stockPrices); i++ {
		ch := stockPrices[i]
		pre := stockPrices[i-1]
		ndy, ndx := ch[1]-pre[1], ch[0]-pre[0]
		if ndy*preDx != ndx*preDY {
			ans++
			preDY, preDx = ndy, ndx

		}
	}
	return ans
}
