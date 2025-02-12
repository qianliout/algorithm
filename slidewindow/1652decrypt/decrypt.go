package main

func main() {

}

func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k == 0 {
		return ans
	}
	sum := make([]int, n*2+1)
	for i := 0; i < 2*n; i++ {
		sum[i+1] = sum[i] + code[i%n]
	}
	for i := 0; i < n; i++ {
		le := i + 1
		if k < 0 {
			le = i + n
		}
		ri := le + k
		if k < 0 {
			le, ri = ri, le
		}

		ans[i] = sum[ri] - sum[le]
	}
	return ans
}
