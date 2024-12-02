package main

import (
	"fmt"
)

func main() {
	// fmt.Println(equationsPossible([]string{"c==c", "b==d", "x!=z"}))
	fmt.Println(equationsPossible([]string{"c!=c"}))
}

// 先用uf
func equationsPossible(equations []string) bool {
	uf := &UF{Fa: make(map[string]string)}
	n := len(equations)
	ss := make([]pair, n)
	for i, ch := range equations {
		p := gen(ch)
		ss[i] = p
		if p.eq {
			uf.Union(p.a, p.b)
		}
	}
	for _, ch := range ss {
		if ch.eq && !uf.IsConnected(ch.a, ch.b) {
			return false
		}
		if !ch.eq && (uf.IsConnected(ch.a, ch.b) || (uf.Find(ch.a) == "" || uf.Find(ch.b) == "")) {
			return false
		}
	}

	return true
}

type pair struct {
	a, b string
	eq   bool
}

func gen(a string) pair {
	ans := pair{}
	if a[1] == '!' {
		ans.a = a[0:1]
		ans.b = a[3:4]
		ans.eq = false
	} else {
		ans.a = a[0:1]
		ans.b = a[3:4]
		ans.eq = true
	}

	return ans
}

type UF struct {
	Fa map[string]string
}

func (u *UF) Find(a string) string {
	if u.Fa[a] != a {
		u.Fa[a] = u.Find(u.Fa[a])
	}
	return u.Fa[a]
}

func (u *UF) Union(a, b string) {
	if u.Fa[a] == "" {
		u.Fa[a] = a
	}
	if u.Fa[b] == "" {
		u.Fa[b] = b
	}

	ar := u.Find(a)
	br := u.Find(b)
	u.Fa[ar] = br
}

func (u *UF) IsConnected(a, b string) bool {
	return u.Find(a) == u.Find(b)
}
