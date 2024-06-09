package main

func main() {

}

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	if sx == fx && sy == fy && t == 1 {
		return false
	}
	if max(abs(sx-fx), abs(sy-fy)) <= t {
		return true
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
