package main

func main() {

}

func mySqrt(x int) int {
	le, ri := 1, x+1
	for le < ri {
		mid := le + (ri-le+1)/2
		if mid >= 1 && mid*mid <= x {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	if le*le <= x {
		return le
	}
	return 0
}
