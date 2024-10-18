package main

import (
	"fmt"
)

func main() {
	fmt.Println(beautifulArray(4))
}

func beautifulArray(n int) []int {
	mem := make(map[int][]int)
	mem[1] = []int{1}
	var dfs func(i int) []int
	dfs = func(i int) []int {
		if _, ok := mem[i]; !ok {
			res := make([]int, 0)
			res1 := dfs((i + 1) / 2)
			for _, ch := range res1 {
				res = append(res, ch*2-1)
			}
			res2 := dfs(i / 2)
			for _, ch := range res2 {
				res = append(res, ch*2)
			}
			mem[i] = res
			return res
		}
		return mem[i]
	}
	res := dfs(n)
	return res
}
