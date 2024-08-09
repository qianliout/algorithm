package main

func main() {

}

func partitionString(s string) int {
	ans, n := 0, len(s)
	cnt := make(map[byte]int)

	for i := 0; i < n; i++ {
		b := byte(s[i])
		if cnt[b] == 0 {
			cnt[b]++
		} else {
			ans++
			cnt = make(map[byte]int)
			cnt[b]++
		}
	}
	if len(cnt) > 0 {
		ans++
	}
	return ans
}
