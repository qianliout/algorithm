package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(groupStrings([]string{"a", "b", "ab", "cde"}))
	// fmt.Println(groupStrings([]string{"a", "ab", "abc"}))
	fmt.Println(groupStrings([]string{"xo", "t", "uhc", "gf"}))
}

func groupStrings(words []string) []int {
	n := len(words)
	uf := NewSizeUnionFind()
	uf.Words = words
	ccWord := make(map[int][]int, n)
	for _, w := range words {
		cc := gen(w)
		ccWord[cc[0]] = gen(w)
	}

	for _, cc := range ccWord {
		for _, ch := range cc {
			if _, ok := uf.Parent[ch]; ok {
				continue
			}
			uf.Parent[ch] = ch
		}
	}
	for _, cc := range ccWord {
		for i := 1; i < len(cc); i++ {
			uf.Union(cc[0], cc[i])
		}
	}

	count := make(map[int]int)
	for _, v := range uf.Parent {
		count[v]++
	}
	group := make(map[int]int)

	for c, ch := range ccWord {
		cc := min(c, ch[0])
		group[uf.Parent[cc]]++
	}
	mx := 0
	for _, v := range group {
		mx = max(mx, v)
	}

	return []int{len(count), mx}
}

func gen(word string) []int {
	ans := make([]int, 0)
	ss := []byte(word)
	ans = append(ans, cal1(ss, -1))
	// 替换
	for i := 0; i < len(ss); i++ {
		pre := ss[i]
		for j := 'a'; j <= 'z'; j++ {
			if byte(j) == pre {
				continue
			}
			ss[i] = byte(j)
			c := cal1(ss, -1)
			if c > 0 {
				ans = append(ans, c)
			}
		}
		ss[i] = pre
	}
	// 删除
	for i := 0; i < len(ss); i++ {
		c := cal1(ss, i)
		if c > 0 {
			ans = append(ans, c)
		}
	}
	// 增加
	ss = append(ss, 'a')
	for j := 'a'; j <= 'z'; j++ {
		ss[len(ss)-1] = byte(j)
		c := cal1(ss, -1)
		ans = append(ans, c)
	}
	sort.Ints(ans)

	return ans
}

func cal1(ss []byte, di int) int {
	ans := 0
	for i, ch := range ss {
		if di != -1 && di == i {
			continue
		}
		ans |= 1 << (int(ch) - int('a'))
	}
	return ans
}

type SizeUnionFind struct {
	Words  []string
	Parent map[int]int // 可以理解成下标i的最终父节点就是 Parent[i]
}

func NewSizeUnionFind() *SizeUnionFind {
	p := make(map[int]int)
	// s := make(map[int]int)
	return &SizeUnionFind{Parent: p}
}

func (u *SizeUnionFind) Find(x int) int {
	if u.Parent[x] != x {
		// 路径压缩
		u.Parent[x] = u.Find(u.Parent[x])
		// 如果在查询中不修改数据
	}
	return u.Parent[x]
}

func (u *SizeUnionFind) Union(x, y int) {
	xRoot := u.Find(x)
	yRoot := u.Find(y)
	if xRoot != yRoot {
		if xRoot <= yRoot {
			u.Union(y, x)
		} else {
			u.Parent[xRoot] = u.Parent[yRoot]
		}
	}
}

func (u *SizeUnionFind) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
