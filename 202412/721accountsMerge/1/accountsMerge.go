package main

import (
	"sort"
)

func main() {

}

func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	ans := make([][]string, 0)
	visit := make(map[string]bool)
	var dfs func(start string) []string
	g := Build(accounts)
	dfs = func(start string) []string {
		if visit[start] {
			return nil
		}
		res := []string{start}
		visit[start] = true
		for _, nx := range g[start] {
			if !visit[nx] {
				res = append(res, dfs(nx)...)
			}
		}
		return res
	}
	// gen ans
	for i := 0; i < n; i++ {
		for j := 1; j < len(accounts[i]); j++ {
			a := dfs(accounts[i][j])
			if len(a) > 0 {
				b := dup(a)
				sort.Strings(b)
				ans = append(ans, append([]string{accounts[i][0]}, b...))
			}
		}
	}
	return ans
}

func Build(accounts [][]string) map[string][]string {
	g := make(map[string][]string)
	for _, acc := range accounts {
		for i := 1; i < len(acc); i++ {
			g[acc[i]] = append(g[acc[i]], acc[1:]...)
		}
	}
	return g
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
	// sort.Strings(ans)
	return ans
}
