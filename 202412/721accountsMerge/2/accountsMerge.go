package main

import (
	"sort"
)

func main() {

}

func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	ac := make([]*Account, n)
	for i, ch := range accounts {
		ac[i] = Gen(ch)
	}
	uf := NewUF(n)
	for i := range ac {
		for j := range ac {
			if i == j {
				continue
			}
			if Same(ac[i], ac[j]) {
				uf.Union(i, j)
			}
		}
	}
	// 输出结果
	ans := make(map[int]*Account)
	for i := 0; i < n; i++ {
		a := uf.Find(i)
		b := ac[i]
		if ans[a] == nil {
			ans[a] = b
		} else {
			for k := range b.Email {
				ans[a].Email[k] = true
			}
		}
	}
	res := make([][]string, 0)
	for _, v := range ans {
		a := make([]string, 0)
		a = append(a, v.Name)

		e := make([]string, 0)
		for k := range v.Email {
			e = append(e, k)
		}
		sort.Strings(e)
		a = append(a, e...)
		res = append(res, a)
	}
	return res
}

type Account struct {
	Name  string
	Email map[string]bool
}

func NewUF(n int) *UF {
	u := &UF{Pa: make([]int, n)}
	for i := 0; i < n; i++ {
		u.Pa[i] = i
	}
	return u
}

type UF struct {
	Pa []int
}

func (u *UF) Union(a, b int) {
	ar := u.Find(a)
	br := u.Find(b)

	// 先不管层级等信息
	u.Pa[ar] = br
}

func (u *UF) Find(a int) int {
	if u.Pa[a] != a {
		u.Pa[a] = u.Find(u.Pa[a])
	}
	return u.Pa[a]
}

func Gen(ac []string) *Account {
	a := &Account{Name: ac[0], Email: make(map[string]bool)}
	for i := 1; i < len(ac); i++ {
		a.Email[ac[i]] = true
	}
	return a
}

func Same(a, b *Account) bool {
	for k := range a.Email {
		if b.Email[k] {
			return true
		}
	}
	return false
}
