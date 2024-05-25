package main

import (
	"fmt"
)

func main() {
	// fmt.Println(countPairs(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	fmt.Println(countPairs(7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}))

	fmt.Println(countPairs(100000, [][]int{{0, 1}}))
	fmt.Println(countPairs(100000, [][]int{}))
}

func countPairs(n int, edges [][]int) int64 {
	in := make([][]int, n)
	for _, ch := range edges {
		in[ch[0]] = append(in[ch[0]], ch[1])
		in[ch[1]] = append(in[ch[1]], ch[0])
	}

	vis := make([]bool, n)
	total := 0
	var ans int64
	for k, v := range vis {
		if v {
			continue
		}
		find := dfs(in, k, vis)
		ans += int64(find) * int64(total)
		total += find
	}

	return int64(ans)
}

func dfs(edges [][]int, start int, visit []bool) int {

	visit[start] = true
	ans := 1
	for _, ch := range edges[start] {
		if visit[ch] {
			continue
		}
		visit[ch] = true
		ans += dfs(edges, ch, visit)
	}
	return ans
}

// time out,超时的原因是每次都 append
func dfs1(edges [][]int, start int, visit []bool) []int {
	ans := make([]int, 0)
	ans = append(ans, start)
	visit[start] = true
	nex := edges[start]
	for _, ch := range nex {
		if visit[ch] {
			continue
		}
		visit[ch] = true
		ans = append(ans, dfs1(edges, ch, visit)...)
	}
	return ans
}
