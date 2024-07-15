package main

func main() {

}

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	p := make([][]int, n)
	for i := range p {
		p[i] = make([]int, n)
	}
	// i和j的 亲密度
	for x, ch := range preferences {
		for score, y := range ch {
			p[x][y] = n - score
		}
	}
	ans := 0
	m := len(pairs)
	for i, ch := range pairs {
		x, y := ch[0], ch[1]
		xOK, yOK := false, false
		for j := 0; j < m; j++ {
			if i == j {
				continue
			}
			u, v := pairs[j][0], pairs[j][1]
			if !xOK && check(p, x, y, u, v) {
				xOK = true
			}
			if !yOK && check(p, y, x, u, v) {
				yOK = true
			}
			if xOK && yOK {
				break
			}
		}
		if xOK {
			ans++
		}
		if yOK {
			ans++
		}
	}
	return ans
}

func check(g [][]int, x, y, u, v int) bool {
	xmp := g[x]
	ump := g[u]
	vmp := g[v]
	if xmp[u] > xmp[y] && ump[x] > ump[v] {
		return true
	}
	if xmp[v] > xmp[y] && vmp[x] > vmp[u] {
		return true
	}

	return false
}
