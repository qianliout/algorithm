package main

func main() {

}

func longestPalindrome(s string) string {
	n := len(s)
	mx := ""
	for i := 0; i < n; i++ {
		a := expand(s, i, i)
		if len(a) > len(mx) {
			mx = a
		}
		if i+1 < n {
			b := expand(s, i, i+1)
			if len(b) > len(mx) {
				mx = b
			}
		}
	}
	return mx
}
func expand(s string, left, right int) string {
	n := len(s)
	for left >= 0 && right < n && s[left] == s[right] {
		left--
		right++
	}

	return s[left+1 : right]
}
