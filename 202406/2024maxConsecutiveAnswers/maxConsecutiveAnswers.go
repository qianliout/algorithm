package main

func main() {

}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(getCnt(answerKey, 'T', k), getCnt(answerKey, 'F', k))
}

func getCnt(s string, c byte, k int) int {
	ans, n := 0, len(s)
	le, ri := 0, 0
	cnt := 0
	for le <= ri && ri < n {
		if s[ri] != c {
			cnt++
		}
		for cnt > k {
			if s[le] != c {
				cnt--
			}
			le++
		}
		ri++
		ans = max(ans, ri-le)
	}
	return ans
}
