package main

func main() {

}

func countGoodSubstrings(s string) int {
	ss := []byte(s)
	n := len(s)
	ans := 0
	for i := 0; i <= n-3; i++ {
		if check(ss[i : i+3]) {
			ans++
		}
	}
	return ans
}

func check(s []byte) bool {
	cnt := make(map[byte]int)
	for _, ch := range s {
		cnt[ch]++
	}
	return len(cnt) == len(s)
}
