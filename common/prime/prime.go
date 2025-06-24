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

// EulerPrime 欧拉筛法（线性筛法）- 时间复杂度O(n)
// 相比埃拉托斯特尼筛法，欧拉筛法保证每个合数只被它的最小质因子筛掉一次
// 这样避免了重复筛选，达到线性时间复杂度
func EulerPrime(n int) []int {
	// notPrime[i] 表示 i 是否为合数（非质数）
	notPrime := make([]bool, n+1)

	// 存储已找到的质数，按从小到大的顺序
	prime := make([]int, 0)

	// 从2开始遍历每个数
	for i := 2; i <= n; i++ {
		// 如果 i 没有被标记为合数，说明 i 是质数
		if !notPrime[i] {
			prime = append(prime, i)
		}

		// 用当前已知的所有质数去筛选合数
		// 关键思想：每个合数只被它的最小质因子筛掉
		for _, j := range prime {
			// 如果 i*j 超过了范围，直接跳出
			if i*j > n {
				break
			}

			// 将 i*j 标记为合数
			// 这里 j 是质数，i*j 的最小质因子就是 j
			notPrime[i*j] = true

			// 关键优化：如果 i 能被 j 整除，说明 j 是 i 的最小质因子
			if i%j == 0 {
				// 此时如果继续用更大的质数 k（k > j）去计算 i*k，
				// 那么 i*k 的最小质因子仍然是 j，而不是 k
				// 这意味着 i*k 应该在处理某个数 m（m < i，且 m*j = i*k）时被筛掉
				// 为了保证每个合数只被最小质因子筛掉一次，这里要 break
				break
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

// Comb 使用动态规划计算组合数 C(n,k) - 时间复杂度O(n*k)，空间复杂度O(n*k)
// 基于帕斯卡三角形的递推关系：C(n,k) = C(n-1,k-1) + C(n-1,k)
// 返回一个二维数组，ans[i][j] 表示 C(i,j) 的值
func Comb(n, k int, mod int) [][]int {
	// 创建二维数组存储组合数结果
	// ans[i][j] 表示从 i 个元素中选择 j 个元素的组合数
	ans := make([][]int, n+1)
	for i := range ans {
		ans[i] = make([]int, k+1)
	}

	// 边界条件：C(0,0) = 1（从0个元素中选0个元素只有1种方法）
	ans[0][0] = 1

	// 填充组合数表
	for i := 1; i <= n; i++ {
		// 边界条件：C(i,0) = 1（从任意i个元素中选0个元素只有1种方法）
		ans[i][0] = 1

		// 计算 C(i,j)，其中 1 <= j <= min(k, i)
		// j <= i 是因为从i个元素中不能选择超过i个元素
		for j := 1; j <= k && j <= i; j++ {
			// 帕斯卡三角形递推公式：C(i,j) = C(i-1,j-1) + C(i-1,j)
			// 含义：从i个元素中选j个 = 包含第i个元素的选法 + 不包含第i个元素的选法
			// C(i-1,j-1)：包含第i个元素，从前i-1个元素中再选j-1个
			// C(i-1,j)：不包含第i个元素，从前i-1个元素中选j个
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]

			// 取模防止整数溢出，在大数计算中非常重要
			ans[i][j] = ans[i][j] % mod
		}
	}
	return ans
}
