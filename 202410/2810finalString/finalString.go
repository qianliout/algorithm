package main

func main() {

}

func finalString(s string) string {
	ans := make([]byte, 0)
	ss := []byte(s)
	for _, ch := range ss {
		if ch != 'i' {
			ans = append(ans, ch)
		} else {
			l, r := 0, len(ans)
			for l < r {
				ans[l], ans[r] = ans[r], ans[l]
				l++
				r--
			}
		}
	}
	return string(ans)
}
