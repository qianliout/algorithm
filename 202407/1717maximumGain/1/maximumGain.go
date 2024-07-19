package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumGain("cdbcbbaaabab", 4, 5))
}

// 不能得到正确的结果，因为只执行了一次
func maximumGain(s string, x int, y int) int {
	ss := []byte(s)
	var dfs func(i int) int
	n := len(s)
	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}
	dfs = func(i int) int {
		if i <= 0 { // 至少要删除两个才得分
			return 0
		}
		if mem[i] != -1 {
			return mem[i]
		}
		res := dfs(i - 1)

		if ss[i] == 'b' && ss[i-1] == 'a' {
			res = max(res, dfs(i-2)+x)
		}
		if ss[i] == 'a' && ss[i-1] == 'b' {
			res = max(res, dfs(i-2)+y)
		}
		mem[i] = res
		return res
	}
	res := dfs(n - 1)
	return res
}
