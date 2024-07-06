package main

func main() {

}

func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	n := len(points)

	x0, y0 := points[0][0], points[0][1]
	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		ans += max(abs(x-x0), abs(y-y0))
		x0, y0 = x, y
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
