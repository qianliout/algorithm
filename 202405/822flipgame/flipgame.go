package main

import (
	"math"
)

func main() {

}

/*
假设给你两个数组A,B。A,B的长度是相同的，你可以任意更换相同下标的两个数，这个时候A,B的数组就会发生变化，
请问你是否能在B数组中找到一个最小值，这个最小值不在A数组中出现，如果存在请返回这个最小值，如果不存在请返回0。
*/

/*
根据题意，只要有数字满足 fronts[i]=backs[i]那么 fronts[i] 绝对不可能是答案，否则 fronts[i]或者 backs[i]作为背面的数字可以满足要求，取最小值作为答案。

我们可以先遍历一遍数组，找到满足 fronts[i]=backs[i] 的数字，存入哈希表 forbidden 中。然后再次遍历数组，找到不在 forbidden中的数字，取最小值作为答案。如果所有数字都在 forbidden中，返回 0。
*/
func flipgame(fronts []int, backs []int) int {
	forbidden := make(map[int]int)
	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] {
			forbidden[fronts[i]]++
		}
	}
	ans := math.MaxInt
	for i, x := range fronts {
		if forbidden[x] == 0 {
			ans = min(ans, x)
		}
		if forbidden[backs[i]] == 0 {
			ans = min(ans, backs[i])
		}
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}
