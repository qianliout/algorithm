package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star", "star", "arts"}))
	fmt.Println(numSimilarGroups([]string{"qihcochwmglyiggvsqqfgjjxu", "gcgqxiysqfqugmjgwclhjhovi",
		"gjhoggxvcqlcsyifmqgqujwhi", "wqoijxciuqlyghcvjhgsqfmgg", "qshcoghwmglygqgviiqfjcjxu",
		"jgcxqfqhuyimjglgihvcqsgow", "qshcoghwmggylqgviiqfjcjxu", "wcoijxqiuqlyghcvjhgsqgmgf",
		"qshcoghwmglyiqgvigqfjcjxu", "qgsjggjuiyihlqcxfovchqmwg", "wcoijxjiuqlyghcvqhgsqgmgf",
		"sijgumvhqwqioclcggxgyhfjq", "lhogcgfqqihjuqsyicxgwmvgj", "ijhoggxvcqlcsygfmqgqujwhi",
		"qshcojhwmglyiqgvigqfgcjxu", "wcoijxqiuqlyghcvjhgsqfmgg", "qshcojhwmglyiggviqqfgcjxu",
		"lhogcgqqfihjuqsyicxgwmvgj", "xscjjyfiuglqigmgqwqghcvho", "lhggcgfqqihjuqsyicxgwmvoj",
		"lhgocgfqqihjuqsyicxgwmvgj", "qihcojhwmglyiggvsqqfgcjxu", "ojjycmqshgglwicfqguxvihgq",
		"sijvumghqwqioclcggxgyhfjq", "gglhhifwvqgqcoyumcgjjisqx"}))
}

func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := NewUf(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if similar(strs[i], strs[j]) {
				uf.Union(i, j)
			}
		}
	}
	return uf.Count
}

func similar(a, b string) bool {
	if a == b {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	n := len(a)
	cnt := 0
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt == 2 || cnt == 0
}

type UF struct {
	Fa    []int
	Count int
}

func (u *UF) Find(a int) int {
	if u.Fa[a] != a {
		u.Fa[a] = u.Find(u.Fa[a])
	}
	return u.Fa[a]
}

func (u *UF) Union(a, b int) {

	ar := u.Find(a)
	br := u.Find(b)
	if ar != br {
		u.Fa[ar] = br
		u.Count--
	}
}

func (u *UF) IsConnected(a, b int) bool {
	return u.Find(a) == u.Find(b)
}

func NewUf(n int) *UF {
	u := &UF{
		Fa:    make([]int, n),
		Count: n,
	}
	for i := 0; i < n; i++ {
		u.Fa[i] = i
	}
	return u
}
