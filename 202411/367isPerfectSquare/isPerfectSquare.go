package main

func main() {

}

func isPerfectSquare(num int) bool {
	le, ri := 1, num+1
	for le < ri {
		mid := le + (ri-le+1)/2
		if mid >= 1 && mid <= num && mid*mid <= num {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	if le*le == num {
		return true
	}
	return false
}
