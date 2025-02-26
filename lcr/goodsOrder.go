package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(goodsOrder("aagew"))
}

func goodsOrder(goods string) []string {
	gg := []byte(goods)
	sort.Slice(gg, func(i, j int) bool { return gg[i] < gg[j] })
	ans := make([]string, 0)
	n := len(goods)
	var dfs func(used []bool, path []byte)
	// 枚举选那一个
	dfs = func(used []bool, path []byte) {
		if len(path) == n {
			ans = append(ans, string(path))
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			// 有重复的时候剪枝
			if i > 0 && gg[i] == gg[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, gg[i])
			dfs(used, path)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	used := make([]bool, n)
	path := make([]byte, 0)
	dfs(used, path)
	return ans
}

// 选和不选的思想不好做
