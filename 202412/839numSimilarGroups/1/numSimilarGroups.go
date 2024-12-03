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
	uf := &UF{Fa: make(map[string]string), Count: n}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if similar(strs[i], strs[j]) {
				uf.Union(strs[i], strs[j])
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
	Fa    map[string]string
	Count int
}

func (u *UF) Find(a string) string {
	if u.Fa[a] != a {
		u.Fa[a] = u.Find(u.Fa[a])
	}
	return u.Fa[a]
}

// 这样写是不行的:["abc","abc"]
func (u *UF) Union(a, b string) {
	if u.Fa[a] == "" {
		u.Fa[a] = a
	}
	if u.Fa[b] == "" {
		u.Fa[b] = b
	}

	ar := u.Find(a)
	br := u.Find(b)
	if ar != br {
		u.Fa[ar] = br
		u.Count--
	}
}

func (u *UF) IsConnected(a, b string) bool {
	return u.Find(a) == u.Find(b)
}
