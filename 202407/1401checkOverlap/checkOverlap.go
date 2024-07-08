package main

func main() {

}

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	x := f(x1, x2, xCenter)
	y := f(y1, y2, yCenter)
	return x*x+y*y <= radius*radius
}

func f(i, j, k int) int {
	if i <= k && k <= j {
		return 0
	}
	return min(abs(i-k), abs(j-k))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
