package main

func main() {

}

func dismantlingAction(arr string) int {
	gg := []byte(arr)
	n := len(gg)
	wind := make(map[byte]int)
	le, ri := 0, 0
	ans := 0
	for le <= ri && ri < n {
		ch := gg[ri]
		wind[ch]++
		ri++
		for wind[ch] > 1 {
			wind[gg[le]]--
			le++
		}
		ans = max(ans, ri-le)
	}
	return ans
}
