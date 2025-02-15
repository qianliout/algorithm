package main

func main() {

}

func countPoints(points [][]int, queries [][]int) []int {
	n := len(queries)
	qq := make([]circle, n)
	pp := make([]point, len(points))
	ans := make([]int, n)
	for i, ch := range queries {
		c := circle{x: ch[0], y: ch[1], r: ch[2]}
		qq[i] = c
	}
	for i, ch := range points {
		pp[i] = point{x: ch[0], y: ch[1]}
	}
	for i, c := range qq {
		cnt := 0
		for _, p := range pp {
			if in(c, p) {
				cnt++
			}
		}
		ans[i] = cnt
	}
	return ans
}

type circle struct {
	x, y, r int
}

type point struct {
	x, y int
}

func in(c circle, p point) bool {
	// 因为是乘法，可以不用abs
	// a := abs(c.x-p.x)*abs(c.x-p.x) + abs(c.y-p.y)*abs(c.y-p.y)
	a := (c.x-p.x)*(c.x-p.x) + (c.y-p.y)*(c.y-p.y)
	b := c.r * c.r
	return a <= b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
