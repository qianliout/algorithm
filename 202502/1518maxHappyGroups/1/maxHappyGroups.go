package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxHappyGroups(9, []int{1, 8, 1, 8, 1, 8, 1, 8, 2, 7, 2, 7, 2, 7, 2, 7, 3, 6, 3, 6, 3, 6, 3, 6, 4, 5, 4, 5, 4, 6}))
}

func maxHappyGroups(m int, groups []int) int {
	cnt := make([]int, m)
	for _, x := range groups {
		cnt[x%m]++
	}

	cache := make(map[string]int)

	var dfs func(left int, cnt []int) int
	dfs = func(left int, cnt []int) int {
		// 会超时
		key := fmt.Sprintf("%d-%v", left, cnt) // 生成唯一的缓存键
		if val, ok := cache[key]; ok {
			return val
		}
		res := 0
		for x, c := range cnt {
			if c > 0 {
				cnt[x]--
				res = max(res, btoi(left == 0)+dfs((left+x)%m, cnt))
				cnt[x]++
			}
		}

		cache[key] = res
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
