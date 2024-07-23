package main

func main() {

}

func areAlmostEqual(s1 string, s2 string) bool {
	a, b := -1, -1
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if a == -1 {
				a = i
				continue
			}
			if b == -1 {
				b = i
				continue
			}
			return false
		}
	}
	if a == -1 {
		return true
	}
	if b == -1 {
		return false
	}
	return s2[a] == s1[b] && s1[a] == s2[b]
}
