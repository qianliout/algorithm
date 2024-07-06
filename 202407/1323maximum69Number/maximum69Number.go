package main

func main() {

}

func maximum69Number(num int) int {
	ans := make([]int, 0)
	n := num
	for n > 0 {
		ans = append(ans, n%10)
		n = n / 10
	}
	le, ri := 0, len(ans)-1
	for le < ri {
		ans[le], ans[ri] = ans[ri], ans[le]
		le++
		ri--
	}
	fir := -1
	for i := 0; i < len(ans); i++ {
		if ans[i] == 6 {
			fir = i
			break
		}
	}
	if fir != -1 {
		ans[fir] = 9
	}
	res := 0
	for _, ch := range ans {
		res = res*10 + ch
	}
	return res
}
