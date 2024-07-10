package main

func main() {

}
func reformat(s string) string {
	a := make([]byte, 0)
	b := make([]byte, 0)
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			a = append(a, byte(ch))
		} else {
			b = append(b, byte(ch))
		}
	}
	if len(a) < len(b) {
		a, b = b, a
	}

	if abs(len(a)-len(b)) > 1 {
		return ""
	}
	ans := make([]byte, len(s))
	start := 0
	for i := 0; i < len(a); i++ {
		ans[start] = a[i]
		start += 2
	}
	start = 1
	for i := 0; i < len(b); i++ {
		ans[start] = b[i]
		start += 2
	}

	return string(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
