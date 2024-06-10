package main

func main() {

}

func longestSemiRepetitiveSubstring(s string) int {
	ss := []byte(s)
	wind := make([]byte, 0)
	ans, ri, n := 0, 0, len(ss)
	for ri < n {
		wind = append(wind, ss[ri])
		ri++
		for !check(wind) {
			wind = wind[1:]
		}
		ans = max(ans, len(wind))
	}
	return ans
}

func check(data []byte) bool {
	cnt := 0
	for i := 0; i < len(data)-1; i++ {
		if data[i] == data[i+1] {
			cnt++
		}
	}
	return cnt <= 1
}
