package prime

// FactorPrime 获取一个数的所有质因子
func FactorPrime(n int) []int {
	var factors []int
	for i := 2; i*i <= n; i++ {
		// 如果 i 能够整除 N，说明 i 为 N 的一个质因子
		if n%i == 0 {
			for n%i == 0 {
				n = n / i
			}
		}
		factors = append(factors, i)
	}
	// 说明再经过操作之后 N 留下了一个素数
	if n != 1 {
		factors = append(factors, n)
	}

	return factors
}

// Eratosthenes 埃筛法
func Eratosthenes(n int) []int {
	isPrime := make([]bool, n+1)
	prime := make([]int, 0)
	for i := 2; i <= n; i++ {
		isPrime[i] = true // 先把2之后的设置成 true
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			prime = append(prime, i)
			// 如果一个数是质数，那么他的所有倍数就都不是质数
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return prime
}

func Eratosthenes2(n int) []int {
	isPrime := make([]bool, n+1)
	prime := make([]int, 0)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	// 一点优化，只计算到 sqrt(n)
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// 因为只计算到sqrt(n),如果只在这里加入答案的话，是加不完的
			// prime = append(prime, i)
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	// 需要在最后全部加入答案
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			prime = append(prime, i)
		}
	}

	return prime
}

// EulerPrime 欧拉筛法，也是一种线性筛法
func EulerPrime(n int) []int {
	notPrime := make([]bool, n+1)

	prime := make([]int, 0)
	for i := 2; i <= n; i++ {
		if !notPrime[i] {
			prime = append(prime, i)
		}
		for _, j := range prime {
			if i*j > n {
				break
			}
			notPrime[i*j] = true
			if i%j == 0 {
				// i % j == 0
				// 换言之，i 之前被 j 筛过了
				// 由于 pri 里面质数是从小到大的，所以 i 乘上其他的质数的结果一定会被
				// j 的倍数筛掉，就不需要在这里先筛一次，所以这里直接 break
				// 掉就好了
			}
		}
	}
	return prime
}

// CalPrime 求每个质数的个数
// 求一个数中某一个质因子的个数
// 比如12：他有2你2和1个3
func CalPrime(n int) map[int]int {
	ks := make(map[int]int)
	for p := 2; p*p <= n; p++ {
		// ks[p] = 1
		if n%p == 0 {
			// k := 0
			for n%p == 0 {
				ks[p]++
				// k++
				n = n / p
			}
			// ks[p] = k
		}
	}
	if n > 1 {
		ks[n] = 1
	}
	return ks
}

// Comb 求组合数C(n,m) 他可以用递推得到 C(n,m) = C(n-1,m-1) + C(n-1,m)
func Comb(n, k int, mod int) [][]int {
	ans := make([][]int, n+1)
	for i := range ans {
		ans[i] = make([]int, k+1)
	}
	// 初值
	ans[0][0] = 1
	for i := 1; i <= n; i++ {
		ans[i][0] = 1
		for j := 1; j <= k && j <= i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
			// 这里一定要取模，不然就会有溢出
			ans[i][j] = ans[i][j] % mod
		}
	}
	return ans
}
