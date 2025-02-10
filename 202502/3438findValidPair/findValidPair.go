package main

func main() {

}

func findValidPair(s string) string {
	cnt := make([]int, 10)
	for _, ch := range s {
		cnt[ch-'0']++
	}
	n := len(s)
	for i := 1; i < n; i++ {
		a := int(s[i-1] - '0')
		b := int(s[i] - '0')
		if a != b && cnt[a] == a && cnt[b] == b {
			return string(s[i-1]) + string(s[i])
		}
	}
	return ""
}
