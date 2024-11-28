package main

func main() {

}

func arrangeCoins(n int) int {
	le, ri := 1, n+1
	for le < ri {
		mid := le + (ri-le+1)/2
		if (1+mid)*mid/2 <= n {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	return le
}
