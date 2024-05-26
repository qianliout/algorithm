package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findAllPeople(6, [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}, 1))
	fmt.Println(findAllPeople(5, [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}}, 1))
	fmt.Println(findAllPeople(5, [][]int{{1, 4, 3}, {0, 4, 3}}, 3))
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][2] < meetings[j][2] })
	know := map[int]bool{0: true, firstPerson: true}
	start := 0
	for start < len(meetings) {
		// 找到时间相同的会议
		end := start
		for end < len(meetings) {
			if meetings[start][2] != meetings[end][2] {
				break
			}
			end++
		}
		g := make(map[int][]int)

		for i := start; i < end; i++ {
			x, y := meetings[i][0], meetings[i][1]
			g[x] = append(g[x], y)
			g[y] = append(g[y], x)
		}
		visit := make(map[int]bool)
		for x := range g {
			if know[x] {
				dfs(g, x, visit, know)
			}
		}
		start = end
	}
	ans := make([]int, 0)
	for i, v := range know {
		if v {
			ans = append(ans, i)
		}
	}
	sort.Ints(ans)
	return ans
}

// 根据知道的专家x,找到他能传播的所有专家
func dfs(g map[int][]int, x int, visit map[int]bool, know map[int]bool) {
	if visit[x] {
		return
	}
	know[x] = true
	next := g[x]
	visit[x] = true
	for _, ch := range next {
		if visit[ch] {
			continue
		}
		dfs(g, ch, visit, know)
	}
}
