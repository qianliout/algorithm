package main

func main() {

}

func countCompleteDayPairs2(hours []int) int64 {
	sub := make([]int, 24)
	for _, ch := range hours {
		sub[ch%24]++
	}
	ans := 0
	for _, ch := range hours {
		c := ch % 24
		ne := (24 - c) % 24
		x := sub[ne]
		if c == ne {
			x--
		}
		ans += x
		if sub[c] > 0 {
			sub[c]--
		}
	}
	return int64(ans)
}

func countCompleteDayPairs(hours []int) int64 {
	sub := make([]int, 24)
	ans := 0
	for _, ch := range hours {
		ch = ch % 24
		next := (24 - ch) % 24
		ans += sub[next]
		sub[ch]++
	}
	return int64(ans)
}
