package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxHappyGroups(9, []int{1, 8, 1, 8, 1, 8, 1, 8, 2, 7, 2, 7, 2, 7, 2, 7, 3, 6, 3, 6, 3, 6, 3, 6, 4, 5, 4, 5, 4, 6}))
}

func maxHappyGroups(m int, groups []int) int {
	// m最大是9
	cnt := [9]int{}

	for _, x := range groups {
		cnt[x%m]++
	}

	cache := make(map[[9]int]int)

	var dfs func(left int, cnt [9]int) int
	dfs = func(left int, cnt [9]int) int {
		if val, ok := cache[cnt]; ok {
			return val
		}
		res := 0
		for x, c := range cnt {

			if c > 0 { // 说明这个 取余 还有组
				cnt[x]--
				res = max(res, btoi(left == 0)+dfs((left+x)%m, cnt))
				// 恢复
				cnt[x]++
			}
		}

		cache[cnt] = res
		return res
	}

	return dfs(0, cnt)
}

// 辅助函数：bool 转 int
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
