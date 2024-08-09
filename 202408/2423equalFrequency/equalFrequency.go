package main

func main() {

}

func equalFrequency(word string) bool {
	f := make([]int, 26)
	for _, ch := range word {
		idx := int(ch) - int('a')
		f[idx]++
	}
	mi, mx := len(word), 0
	for _, ch := range f {
		if ch > 0 {
			mi = min(mi, ch)
			mx = max(mx, ch)
		}
	}
	for _, ch := range word {
		idx := int(ch) - int('a')
		f[idx]--
		if check(f) {
			return true
		}
		f[idx]++
	}
	return false
}

func check(a []int) bool {
	mi, mx := 110, 0
	for _, ch := range a {
		if ch > 0 {
			mi = min(mi, ch)
			mx = max(mx, ch)
		}
	}
	return mx == mi
}
