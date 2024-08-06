package main

func main() {

}

func minSteps(s string, t string) int {
	cnt1, cnt2 := make(map[byte]int), make(map[byte]int)
	key := make(map[byte]int)
	for _, b := range []byte(s) {
		cnt1[b]++
		key[b]++
	}
	for _, b := range []byte(t) {
		cnt2[b]++
		key[b]++
	}
	ans := 0
	for k := range key {
		ans += abs(cnt2[k] - cnt1[k])
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
