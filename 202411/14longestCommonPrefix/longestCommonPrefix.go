package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	mid := n / 2
	left := longestCommonPrefix(strs[:mid])
	right := longestCommonPrefix(strs[mid:])
	return help(left, right)
}

func help(a, b string) string {
	if len(a) > len(b) {
		return help(b, a)
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[:i]
		}
	}
	return a
}
