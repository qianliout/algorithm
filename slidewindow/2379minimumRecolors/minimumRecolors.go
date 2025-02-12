package main

func main() {

}

func minimumRecolors(blocks string, k int) int {
	le, ri, n := 0, 0, len(blocks)
	win, ans := 0, n
	for le <= ri && ri < n {
		win += check(blocks[ri])
		ri++
		if ri-le >= k {
			ans = min(ans, win)
		}
		if ri-le >= k {
			win -= check(blocks[le])
			le++
		}
	}
	return ans
}
func check(b byte) int {
	if b == byte('W') {
		return 1
	}
	return 0
}
