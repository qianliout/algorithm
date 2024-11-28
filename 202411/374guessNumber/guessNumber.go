package main

func main() {

}

func guessNumber(n int) int {
	le, ri := 1, n+1
	for le < ri {
		mid := le + (ri-le)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			ri = mid - 1
		} else if guess(mid) == 1 {
			le = mid + 1
		}
	}
	return 0
}

func guess(num int) int {
	return 0
}

func firstBadVersion(n int) int {
	le, ri := 1, n+1
	for le < ri {
		mid := le + (ri-le)/2
		if isBadVersion(mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}

func isBadVersion(version int) bool {
	return true
}
