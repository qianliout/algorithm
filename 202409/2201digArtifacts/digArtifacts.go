package main

func main() {

}

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	f := make([]bool, n*n)
	for _, d := range dig {
		idx := d[0]*n + d[1]
		f[idx] = true
	}
	ans := 0
	for _, a := range artifacts {
		nu := cal(a, n)
		if check(f, nu) {
			ans++
		}
	}
	return ans
}

func cal(art []int, n int) []int {
	ans := make([]int, 0)
	a, b, c, d := art[0], art[1], art[2], art[3]
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			ans = append(ans, i*n+j)
		}
	}
	return ans
}

func check(f []bool, ans []int) bool {
	for _, d := range ans {
		if !f[d] {
			return false
		}
	}
	return true
}
