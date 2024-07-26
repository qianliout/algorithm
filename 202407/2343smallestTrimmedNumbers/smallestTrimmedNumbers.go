package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestTrimmedNumbers([]string{"102", "473", "251", "814"}, [][]int{{1, 1}}))
	fmt.Println(smallestTrimmedNumbers([]string{"102", "473", "251", "814"}, [][]int{{4, 2}}))
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	ans := make([]int, len(queries))
	n := len(nums)
	for i, ch := range queries {
		ret := cal(nums, ch[1])
		sort.SliceStable(ret, func(i, j int) bool { return ret[i].value < ret[j].value })
		ans[i] = ret[min(n-1, ch[0]-1)].id
	}
	return ans
}

type pari struct {
	id    int
	value string // 不能返回 int，因为会有越界的分险
}

func cal(nums []string, ti int) []pari {
	ans := make([]pari, len(nums))
	for i, ch := range nums {
		n := max(0, len(ch)-ti)
		tex := ch[n:]
		ans[i] = pari{id: i, value: tex}
	}
	return ans
}
