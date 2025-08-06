package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	wind := make(map[byte]int)
	ans := 0
	l := 0
	n := len(s)
	for i := 0; i < n; i++ {
		ch := byte(s[i])
		wind[ch]++
		for wind[ch] > 1 {
			wind[byte(s[l])]--
			l++
		}
		ans = max(ans, i-l+1)
	}
	return ans
}
