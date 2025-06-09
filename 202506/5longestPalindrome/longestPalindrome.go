package main

func main() {

}

func longestPalindrome(s string) string {
	ss := []byte(s)
	ans := 0
	for i := 0; i < len(s); i++ {

	}

	return ans
}

func help(ss []byte, i, j int) string {
	le, ri := i, j
	for le >= 0 && ri < len(ss) {
		if ss[le] == ss[ri] {
			le--
			ri++
		}
	}
	return string(ss[le : ri+1])
}
