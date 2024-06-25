package main

import (
	"fmt"
	"strings"

	. "outback/algorithm/common/unionfind"
)

func main() {
	fmt.Println(equationsPossible([]string{"c==c", "b==d", "x!=z"}))
}

func equationsPossible(equations []string) bool {
	uf := NewRankUnionFind(30)
	for _, ch := range equations {
		if strings.Contains(ch, "==") {
			split := strings.Split(ch, "==")
			a, b := int(split[0][0])-int('a'), int(split[1][0])-int('a')
			uf.Union(a, b)
		}
	}
	for _, ch := range equations {
		if strings.Contains(ch, "!=") {
			split := strings.Split(ch, "!=")
			a, b := int(split[0][0])-int('a'), int(split[1][0])-int('a')
			if uf.IsConnected(a, b) {
				return false
			}
		}
	}
	return true
}

// type RankUnionFind struct {
// 	Parent map[string]string // 可以理解成下标i的最终父节点就是 Parent[i]
// }
//
// func NewRankUnionFind() *RankUnionFind {
// 	p := make(map[string]string)
// 	return &RankUnionFind{Parent: p}
// }
//
// func (u *RankUnionFind) Find(x string) string {
// 	if u.Parent[x] != x {
// 		u.Parent[x] = u.Find(u.Parent[x])
// 	}
// 	return u.Parent[x]
// }
//
// func (u *RankUnionFind) Union(x, y string) {
// 	xRoot := u.Find(x)
// 	yRoot := u.Find(y)
// 	u.Parent[yRoot] = xRoot
// }
//
// func (u *RankUnionFind) IsConnected(x, y string) bool {
// 	return u.Find(x) == u.Find(y)
// }
