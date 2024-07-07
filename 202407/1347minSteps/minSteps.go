package main

func main() {

}

func minSteps(s string, t string) int {
	n := len(s)
	if len(s) != len(t) {
		return 0
	}
	cnt1 := make(map[byte]int)
	cnt2 := make(map[byte]int)
	cnt := make(map[byte]int)
	for i := 0; i < n; i++ {
		cnt1[s[i]]++
		cnt2[t[i]]++
		cnt[s[i]]++
		cnt[t[i]]++
	}
	ans := 0
	for k := range cnt {
		ans += abs(cnt1[k] - cnt2[k])
	}
	return ans / 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
