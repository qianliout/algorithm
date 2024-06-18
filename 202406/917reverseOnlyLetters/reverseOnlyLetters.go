package main

func main() {

}

func reverseOnlyLetters(s string) string {
	ss := []byte(s)
	start, end := 0, len(s)-1
	for start < end {
		if !check2(byte(s[start])) {
			start++
			continue
		}
		if !check2(byte(s[end])) {
			end--
			continue
		}
		ss[start], ss[end] = ss[end], ss[start]
		start++
		end--
	}
	return string(ss)
}

func check2(a byte) bool {
	if a >= 'a' && a <= 'z' {
		return true
	}
	if a >= 'A' && a <= 'Z' {
		return true
	}
	return false
}
