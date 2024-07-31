package main

func main() {

}

func interchangeableRectangles(rectangles [][]int) int64 {
	g := make(map[int]map[int]int)
	for _, ch := range rectangles {
		x, y := ch[0], ch[1]
		c := gcd(x, y)
		x = x / c
		y = y / c
		if g[x] == nil {
			g[x] = make(map[int]int)
		}
		g[x][y]++
	}
	ans := 0
	for _, vv := range g {
		for _, cnt := range vv {
			ans += (cnt * (cnt - 1)) / 2
		}
	}
	return int64(ans)

}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
