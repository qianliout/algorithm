package main

import (
	"fmt"
	"sort"
	// . "outback/algorithm/common/unionfind"
)

func main() {
	fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}))
}

// 会涉及到多次合并，所以这样直接模拟的方式搞不定
func accountsMerge(accounts [][]string) [][]string {
	ans := make([][]string, 0)
	uf := NewSizeUnionFind()

	for i, acc := range accounts {
		name := acc[0]
		for _, em := range acc[1:] {
			uf.Union(em, name)
		}

	}

	return ans
}

type SizeUnionFind struct {
	Parent map[string]string
	Rank   map[string]int
}

func NewSizeUnionFind() *SizeUnionFind {

}

func (vi *SizeUnionFind) Union(x, y string) {

}

func (vi *SizeUnionFind) Find(x string) string {
	if vi.Parent[x] != x {
		vi.Parent[x] = vi.Find(vi.Parent[x])
	}
	return vi.Parent[x]
}

func (vi *SizeUnionFind) IsConnect(x, y string) bool {
	return vi.Find(x) == vi.Find(y)
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
	sort.Strings(ans)
	return ans
}
