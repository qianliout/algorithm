package main

func main() {

}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	uf := &UF{Fa: make([]byte, 26)}
	for i := 'a'; i <= 'z'; i++ {
		uf.Fa[i-'a'] = byte(i)
	}
	n := len(s1)
	for i := 0; i < n; i++ {
		uf.Union(s1[i], s2[i])
	}
	m := len(baseStr)
	ans := make([]byte, m)
	for i := 0; i < m; i++ {
		ans[i] = uf.Find(baseStr[i])
	}
	return string(ans)
}

type UF struct {
	Fa []byte
}

func (u *UF) Find(a byte) byte {
	idx := int(a - 'a')
	if u.Fa[idx] != byte(idx+'a') {
		u.Fa[idx] = u.Find(u.Fa[idx])
	}
	return u.Fa[idx]
}

func (u *UF) Union(a, b byte) {
	ar := u.Find(a)
	br := u.Find(b)

	// 题目中要求字典序最小的
	if ar < br {
		u.Fa[int(br)-'a'] = ar
	} else {
		u.Fa[int(ar)-'a'] = br
	}
}

func (u *UF) IsConnected(a, b byte) bool {
	return u.Find(a) == u.Find(b)
}
