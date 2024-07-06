package main

import "fmt"

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
}

func groupThePeople(groupSizes []int) [][]int {
	g := make(map[int][]int)
	for i, ch := range groupSizes {
		g[ch] = append(g[ch], i)
	}
	ans := make([][]int, 0)
	for k, v := range g {
		if len(v) <= k {
			ans = append(ans, v)
		} else {
			// split
			ans = append(ans, split(v, k)...)
		}

	}
	return ans

}

func split(arr []int, k int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0)
	for _, ch := range arr {
		ans = append(ans, ch)
		if len(ans) >= k {
			res = append(res, append([]int{}, ans...))
			ans = ans[:0]
		}
	}
	return res
}
