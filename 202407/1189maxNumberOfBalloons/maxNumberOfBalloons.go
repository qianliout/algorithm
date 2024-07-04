package main

func main() {

}

// balloon
func maxNumberOfBalloons(text string) int {
	cnt := make(map[byte]int)
	for _, ch := range text {
		cnt[byte(ch)]++
	}
	cnt['l'] /= 2
	cnt['o'] /= 2
	inf := len(text)
	ans := inf
	for _, ch := range []byte("balon") {
		ans = min(ans, cnt[byte(ch)])
	}
	return ans
}
