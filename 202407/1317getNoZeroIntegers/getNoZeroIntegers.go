package main

func main() {

}

func getNoZeroIntegers(n int) []int {
	for i := 1; i <= n/2; i++ {
		a, b := i, n-i
		if check(a) && check(b) {
			return []int{a, b}
		}
	}
	return []int{1, 1}

}

func check(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return false
		}
		n = n / 10
	}

	return true
}
