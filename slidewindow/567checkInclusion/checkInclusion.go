package main

func main() {

}

func checkInclusion(s1 string, s2 string) bool {
	le, ri, n := 0, 0, len(s2)
	a := make([]int, 26)
	for _, ch := range s1 {
		a[int(ch)-int('a')]++
	}
	win := make([]int, 26)
	for le <= ri && ri < n {
		idx := int(s2[ri]) - int('a')
		win[idx]++
		ri++
		if check(a, win) {
			return true
		}
		if ri-le >= len(s1) {
			idx = int(s2[le]) - int('a')
			win[idx]--
			le++
		}
	}
	return false
}

func check(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
