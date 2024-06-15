package main

func main() {

}

func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-2; i++ {
		if !f(n, i) {
			return false
		}
	}
	return true
}

func f(a, n int) bool {
	ans := make([]int, 0)
	for a > 0 {
		ans = append(ans, a%n)
		a = a / n
	}
	if len(ans) == 0 {
		return false
	}
	le, ri := 0, len(ans)-1
	for le < ri {
		if ans[le] != ans[ri] {
			return false
		}
		le++
		ri--
	}
	return true
}
