package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}))
}

func accountsMerge(accounts [][]string) [][]string {
	// 只建立各个 email 之间的关系,通过 email 的关系确定用户的关系
	graph := make(map[string][]string) // email -> related email list
	for _, acc := range accounts {
		for j := 1; j < len(acc); j++ {
			graph[acc[j]] = append(graph[acc[j]], acc[1:]...)
		}
	}
	ans := make([][]string, 0)
	visit := make(map[string]bool)
	for _, acc := range accounts {
		for j := 1; j < len(acc); j++ {
			em := dfs(graph, acc[j], visit)
			if len(em) > 0 {
				ans = append(ans, append([]string{acc[0]}, dup(em)...))
			}
		}
	}

	return ans
}

func dfs(graph map[string][]string, em string, visit map[string]bool) []string {
	ans := make([]string, 0)
	if visit[em] {
		return ans
	}
	ans = append(ans, em)
	visit[em] = true
	for _, ems := range graph[em] {
		if visit[ems] {
			continue
		}
		ans = append(ans, dfs(graph, ems, visit)...)
	}
	return ans
}

func dup(strs []string) []string {
	exit := make(map[string]bool)
	ans := make([]string, 0)
	for _, ch := range strs {
		if !exit[ch] {
			ans = append(ans, ch)
			exit[ch] = true
		}
	}
	// 不好的写法，兼顾了排序
	sort.Strings(ans)
	return ans
}
