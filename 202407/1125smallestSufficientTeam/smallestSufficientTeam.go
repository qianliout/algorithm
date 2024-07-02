package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(smallestSufficientTeam([]string{"java", "nodejs", "reactjs"}, [][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}}))
	fmt.Println(^2)
}

func smallestSufficientTeam(s []string, p [][]string) []int {
	sk1 := make(map[int]string)
	sk2 := make(map[string]int)
	for i, x := range s {
		sk1[i] = x
		sk2[x] = i
	}
	m, n := len(s), len(p)
	mask := make([]int, n)
	for i, ch := range p {
		ski := 0
		for _, c := range ch {
			ski = ski | (1 << sk2[c])
		}
		mask[i] = ski
	}

	var dfs func(i, j int) int
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, 1<<m)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		if j == 0 {
			return 0 // 背包已装满
		}
		if i < 0 { // 说明没有办法选取人员，返回全集
			return 1<<n - 1
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		// 不选 i
		res1 := dfs(i-1, j)

		// 选i
		res2 := dfs(i-1, j&^mask[i]) | (1 << i)

		a := bits.OnesCount(uint(res1))
		b := bits.OnesCount(uint(res2))
		if a <= b {
			mem[i][j] = res1
			return res1
		}
		mem[i][j] = res2
		return res2
	}
	res := dfs(n-1, 1<<len(s)-1)
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if (res>>i)&1 > 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
