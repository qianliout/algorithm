package main

func main() {

}

func breakPalindrome(palindrome string) string {
	ss := []byte(palindrome)
	n := len(palindrome)
	ans := ""
	for i := 0; i < n; i++ {
		for j := 'a'; j <= 'z'; j++ {
			if ss[i] == byte(j) {
				continue
			}
			pre := ss[i]
			ss[i] = byte(j)

			if !check(string(ss)) {
				// 不能直接返回，如果直接返回，这种情况就会出错： abccba
				if ans == "" || ans > string(ss) {
					ans = string(ss)
				}
			}
			ss[i] = pre
		}
	}
	return ans
}

func check(s string) bool {
	le, ri := 0, len(s)-1
	for le < ri {
		if s[le] != s[ri] {
			return false
		}
		le++
		ri--
	}
	return true
}
