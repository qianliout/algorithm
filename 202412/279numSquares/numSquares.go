package main

func main() {

}

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = i
		for j := 0; j*j <= i; j++ {
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}
	return f[n]
}
