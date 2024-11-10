package main

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)
	for i := 0; i < len(s); i++ {
		a := s[i] - 'a'
		b := t[i] - 'a'
		cnt1[a]++
		cnt2[b]++
	}
	for i := range cnt1 {
		if cnt2[i] != cnt1[i] {
			return false
		}
	}
	return true
}
