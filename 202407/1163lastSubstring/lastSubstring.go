package main

func main() {

}

func lastSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}
	mxByte := s[0]
	for _, ch := range s {
		mxByte = max(mxByte, byte(ch))
	}
	n := len(s)

	ans := ""
	for i := n - 1; i >= 0; i-- {
		if s[i] != mxByte {
			continue
		}
		ans = max(ans, s[i:])
	}
	return ans
}
