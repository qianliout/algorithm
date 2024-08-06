package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(sortJumbled([]int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38}))
	fmt.Println(sortJumbled([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func sortJumbled(mapping []int, nums []int) []int {
	n := len(nums)
	pairs := make([]pair, n)
	for i, ch := range nums {
		pairs[i] = pair{
			val: ch,
			nex: cal(mapping, ch),
		}
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i].nex != pairs[j].nex {
			return pairs[i].nex < pairs[j].nex
		} else {
			return i < j
		}
	})
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = pairs[i].val
	}
	return ans
}

func cal(mapping []int, num int) int {
	if num == 0 {
		return mapping[0]
	}
	ans := make([]int, 0)
	for num != 0 {
		ans = append(ans, mapping[num%10])
		num = num / 10
	}
	ret := 0
	for i := len(ans) - 1; i >= 0; i-- {
		ret = ret*10 + ans[i]
	}
	return ret
}

type pair struct {
	val int
	nex int
}
