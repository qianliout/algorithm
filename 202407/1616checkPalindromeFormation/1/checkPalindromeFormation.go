package main

func main() {

}

func checkPalindromeFormation(a string, b string) bool {
	if check(a, b) || check(b, a) {
		return true
	}
	a = reverse(a)
	b = reverse(b)
	if check(a, b) || check(b, a) {
		return true
	}

	return false
}

func reverse(s string) string {
	ss := []byte(s)
	le, ri := 0, len(ss)-1
	for le < ri {
		ss[le], ss[ri] = ss[ri], ss[le]
		le++
		ri--
	}
	return string(ss)
}

func check(a, b string) bool {
	n := len(a)
	flag := true
	for i := 0; i < n/2; i++ {
		if flag {
			if a[i] != b[n-i-1] {
				flag = false
			}
		}
		if !flag {
			if a[i] != a[n-i-1] {
				return false
			}
		}
	}
	return true
}
