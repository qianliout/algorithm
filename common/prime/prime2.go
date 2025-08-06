package prime

func EulerPrime2(n int) []int {

	prim := make([]int, 0)
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true // 默认全是
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			prim = append(prim, i)
			break
		}
		for _, p := range prim {
			if p*i > n {
				break
			}
			isPrime[p*i] = false
			if i%p == 0 {
				break
			}
		}
	}
	return prim
}
