package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSmallestString(3, 27))
}

func getSmallestString(n int, k int) string {
	ans := make([]int, n)
	res := ""
	var dfs func(i int, s int)
	dfs = func(i int, s int) {
		if i >= n || i < 0 {
			if s == 0 {
				lev := gen(ans)
				if (lev != "" && lev < res) || res == "" {
					res = lev
				}
				return
			}
			return
		}

		for j := 1; j <= 26; j++ {
			// 贪心的做法
			if j > s {
				break
			}
			pre := ans[i]
			ans[i] = j

			dfs(i+1, s-j)

			ans[i] = pre
		}
	}
	dfs(0, k)
	return res
}

func gen(nums []int) string {

	ans := make([]byte, len(nums))
	for i, ch := range nums {
		ans[i] = byte(ch - 1 + int('a'))
	}
	return string(ans)
}
