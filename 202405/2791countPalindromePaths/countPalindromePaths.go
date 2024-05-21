package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPalindromePaths([]int{-1, 0, 0, 1, 1, 2}, "acaabc"))
}

func countPalindromePaths(parent []int, s string) int64 {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		ch := parent[i]
		g[ch] = append(g[ch], i)
	}

	cnt := make(map[int]int)
	cnt[0] = 1
	ans := dfs(g, s, 0, 0, cnt)
	return int64(ans)
}

// 有错
func dfs1(g [][]int, s string, start int, sum int, cnt map[int]int, ans *int) {
	for _, w := range g[start] {
		sum = sum ^ (1 << (s[w] - 'a'))
		// 偶数
		*ans += cnt[sum]
		// 奇数
		for i := 0; i < 26; i++ {
			pre := sum ^ (1 << i)
			*ans += cnt[pre]
		}
		cnt[sum]++
		dfs1(g, s, w, sum, cnt, ans)
	}
	return
}

func dfs2(g [][]int, s string, start int, sum int, cnt map[int]int, ans *int) {
	for _, w := range g[start] {
		// 这里必须要重新新开一个变量 x，是为什么呢
		x := sum ^ (1 << (s[w] - 'a'))
		// 偶数
		*ans += cnt[x]
		// 奇数
		for i := 0; i < 26; i++ {
			pre := x ^ (1 << i)
			*ans += cnt[pre]
		}
		cnt[x]++
		dfs2(g, s, w, x, cnt, ans)
	}
	return
}

func dfs(g [][]int, s string, start int, sum int, cnt map[int]int) int {
	ans := 0
	for _, w := range g[start] {
		// 这里必须要重新新开一个变量 x，是为什么呢
		// 是因为同层的循环最初时都是共用一个 sum,如果这里改了，那么下一个 w 的值就不是原来的 sum 值了，所以一定要新建一个变量
		x := sum ^ (1 << (s[w] - 'a'))
		// 偶数
		ans += cnt[x]
		// 奇数
		for i := 0; i < 26; i++ {
			pre := x ^ (1 << i)
			ans += cnt[pre]
		}
		cnt[x]++
		ans += dfs(g, s, w, x, cnt)
	}
	return ans
}

// func dfs(g [][]int, s string, start int, sum int, cnt map[int]int) int {
// 	ans := 0
// 	for _, w := range g[start] {
// 		// 这里必须要重新新开一个变量 x，是为什么呢,没有想明白
// 		sum = sum ^ (1 << (s[w] - 'a'))
// 		fmt.Println("sum", sum)
// 		// 偶数
// 		ans += cnt[sum]
// 		// 奇数
// 		for i := 0; i < 26; i++ {
// 			pre := sum ^ (1 << i)
// 			ans += cnt[pre]
// 		}
// 		cnt[sum]++
// 		ans += dfs(g, s, w, sum, cnt)
// 	}
// 	return ans
// }
