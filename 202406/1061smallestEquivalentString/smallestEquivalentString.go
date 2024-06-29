package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestEquivalentString("parker", "morris", "parser"))
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	uf := UniFond{Parent: make([]int, 26)}
	for i := 0; i < 26; i++ {
		uf.Parent[i] = i
	}
	for i := 0; i < len(s1); i++ {
		id1 := int(s1[i]) - int('a')
		id2 := int(s2[i]) - int('a')
		uf.Union(id1, id2)
	}
	ans := make([]byte, 0)

	for _, ch := range baseStr {
		p := uf.Find(int(ch)-int('a')) + 'a'
		ans = append(ans, byte(p))
	}
	return string(ans)
}

type UniFond struct {
	Parent []int
}

func (vi *UniFond) Find(x int) int {
	if vi.Parent[x] != x {
		vi.Parent[x] = vi.Find(vi.Parent[x])
	}
	return vi.Parent[x]
}

func (vi *UniFond) Union(x, y int) {
	xp := vi.Find(x)
	yp := vi.Find(y)
	// 题目中要求找最小的
	if xp > yp {
		vi.Union(y, x)
	} else if xp < yp {
		// 这里要好好理解
		vi.Parent[yp] = vi.Parent[xp]
	}
}

func (vi *UniFond) Connected(x, y int) bool {
	return vi.Find(x) == vi.Find(y)
}
