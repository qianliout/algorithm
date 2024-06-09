package main

func main() {

}

func countPalindromicSubsequence(s string) int {
	n := len(s)
	res := 0
	for i := 'a'; i <= 'z'; i++ {
		l, r := 0, n-1
		for l < r && int32(s[l]) != i {
			l++
		}
		for r > l && int32(s[r]) != i {
			r--
		}
		if r-l < 2 {
			continue
		}
		exit := make(map[uint8]bool)
		for j := l + 1; j <= r-1; j++ {
			exit[s[j]] = true
		}
		res += len(exit)
	}
	return res
}
