package main

func main() {

}

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	i, n := 0, len(s)
	for i < n {
		if words[i][0] != s[i] {
			return false
		}
		i++
	}
	return true
}
