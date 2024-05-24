package main

func main() {

}

func hasAlternatingBits(n int) bool {
	return start(n, true) || start(n, false)

}

func start(n int, one bool) bool {
	for n != 0 {
		if n&1 == 1 == one {
			return false
		}
		n = n >> 1
		one = !one
	}
	return true
}
