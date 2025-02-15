package main

func main() {

}

func minFlips(s string) int {
	n := len(s)
	cnt := 0
	target := "01"
	for i := 0; i < n; i++ {
		if s[i] != target[i%2] {
			cnt++
		}
	}
	ans := min(cnt, n-cnt)

	for i := 0; i < n; i++ {
		if s[i] != target[i%2] {
			cnt--
		}
		if s[i] != target[(i+n)%2] {
			cnt++
		}
		ans = min(ans, cnt, n-cnt)
	}

	return ans
}
