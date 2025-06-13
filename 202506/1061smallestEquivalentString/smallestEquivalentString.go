package main

func main() {

}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	n := min(len(s1), len(s2))
	uf := NewUF()
	for i := 0; i < n; i++ {
		x, y := int(s1[i])-int('a'), int(s2[i])-int('a')
		uf.Union(x, y)
	}
	ans := make([]byte, 0)
	for _, c := range baseStr {
		y := uf.Find(int(c) - int('a'))
		ans = append(ans, byte(y+'a'))
	}
	return string(ans)
}

type UF struct {
	Parent []int
}

func NewUF() *UF {
	// 字符串s1, s2, and baseStr 仅由从 'a' 到 'z' 的小写英文字母组成。
	pa := make([]int, 26)
	for i := 0; i < 26; i++ {
		pa[i] = i
	}
	return &UF{Parent: pa}
}

func (u *UF) Union(x, y int) {
	xr := u.Find(x)
	yr := u.Find(y)
	if xr == yr {
		return
	}
	if xr < yr {
		u.Parent[yr] = xr
	} else {
		u.Parent[xr] = yr
	}
}

func (u *UF) Find(x int) int {
	if u.Parent[x] != x {
		u.Parent[x] = u.Find(u.Parent[x])
	}
	return u.Parent[x]
}

func (u *UF) Connect(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
