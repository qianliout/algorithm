package main

func main() {

}

func findTheLongestBalancedSubstring(s string) int {
	z, o, ans := 0, 0, 0
	right := 0
	n := len(s)
	for right < n {
		if s[right] == '0' {
			if o > 0 {
				z = 1
				o = 0
			} else {
				z++
			}
		} else {
			o++
			ans = max(ans, 2*min(z, o))
		}
		right++
	}

	return ans
}
