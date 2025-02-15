package main

func main() {

}

func findAnagrams(s2 string, s1 string) []int {
	ans := make([]int, 0)
	le, ri, m, n := 0, 0, len(s1), len(s2)
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
			ans = append(ans, le)
		}
		if ri-le >= m {
			idx = int(s2[le]) - int('a')
			win[idx]--
			le++
		}
	}
	return ans
}

func check(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
