package prime

// FactorPrime 获取一个数的所有质因子
// 例如 n = 12，返回 [2, 3]
// 算法思想：
// 1. 从 2 开始遍历到 sqrt(n)。
// 2. 如果 i 能整除 n，说明 i 是 n 的一个质因子。
// 3. 将 i 加入结果列表，并持续用 i 去除 n，直到 n 不能再被 i 整除，这样可以确保将所有 i 的倍数都从 n 中移除。
// 4. 遍历结束后，如果 n 不为 1，说明 n 本身也是一个质数（大于sqrt(n)的质因子最多只有一个）。
func FactorPrime(n int) []int {
	var factors []int
	// 从最小的质数 2 开始尝试
	for i := 2; i*i <= n; i++ {
		// 如果 i 是 n 的因子
		if n%i == 0 {
			// 将 i 添加到质因子列表
			factors = append(factors, i)
			// 持续用 i 去除 n，以消除所有 i 的倍数
			for n%i == 0 {
				n = n / i
			}
		}
	}
	// 如果最后 n 大于 1，说明 n 本身也是一个质因子
	// 例如 n=10, i=2, factors=[2], n=5. 循环结束. n=5 > 1, factors=[2,5]
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

// Eratosthenes (埃拉托斯特尼筛法)
// 用于找出小于等于 n 的所有质数。
// 算法思想：
//  1. 创建一个布尔数组 isPrime，初始化所有数为质数（true）。
//  2. 从 2 开始遍历到 n。
//  3. 如果当前数 i 是质数 (isPrime[i] is true):
//     a. 将 i 添加到质数列表。
//     b. 将所有 i 的倍数（从 i*i 开始）标记为非质数 (false)。因为小于 i*i 的 i 的倍数（如 2*i, 3*i）必然已经被更小的质数（如 2, 3）筛过了。
//  4. 遍历结束后，质数列表即为所求。
//
// 时间复杂度: O(n log log n)
// 空间复杂度: O(n)
func Eratosthenes(n int) []int {
	// isPrime[i] 表示数字 i 是否是质数
	isPrime := make([]bool, n+1)
	// 存储找到的质数
	prime := make([]int, 0)

	// 初始化：假设从 2 到 n 所有的数都是质数
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= n; i++ {
		// 如果 i 是质数
		if isPrime[i] {
			// 将 i 加入质数列表
			prime = append(prime, i)
			// 将 i 的所有倍数标记为非质数
			// 优化：从 i*i 开始标记，因为小于 i*i 的倍数肯定已经被更小的质数筛过了
			// 例如，当 i=5 时，2*5, 3*5, 4*5 已经被 2 和 3 筛过了
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return prime
}

// Eratosthenes2 是埃氏筛法的优化版本。
// 主要优化点：
//  1. 外层循环只需要遍历到 sqrt(n)，因为任何大于 sqrt(n) 的合数，其最小质因子必然小于或等于 sqrt(n)。
//     所以在处理到 sqrt(n) 时，所有合数都已经被筛掉了。
//  2. 质数的收集被移到了筛法循环之后，单独进行一次遍历来收集所有仍然被标记为质数的数。
func Eratosthenes2(n int) []int {
	isPrime := make([]bool, n+1)
	prime := make([]int, 0)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// 优化：外层循环到 sqrt(n) 即可
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// 这里不把 i 加入 prime 列表，因为外层循环只到 sqrt(n)，会漏掉 (sqrt(n), n] 之间的质数
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// 遍历 isPrime 数组，收集所有质数
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			prime = append(prime, i)
		}
	}

	return prime
}

// EulerPrime (欧拉筛法，也叫线性筛法)
// 相比埃氏筛法，欧拉筛法保证每个合数只被其“最小质因子”筛掉一次，从而达到 O(n) 的时间复杂度。
// 算法核心：
//  1. 遍历 2 到 n 的所有数 i。
//  2. 如果 i 是质数，则加入质数列表。
//  3. 遍历已找到的质数列表 prime。对于每个质数 p in prime:
//     a. 将 i*p 标记为合数。这里 p 就是 i*p 的最小质因子（因为 prime 是有序的）。
//     b. 如果 i % p == 0，则 break。这是关键。
//     - 当 i % p == 0，说明 p 是 i 的最小质因子。
//     - 那么对于任何大于 p 的质数 p'，i*p' 这个合数，它的最小质因子是 p，而不是 p'。
//     - 为了保证“每个合数只被其最小质因子筛掉”，必须在这里停止。
func EulerPrime(n int) []int {
	// notPrime[i] 表示 i 是否为合数（非质数）
	notPrime := make([]bool, n+1)
	// 存储已找到的质数，按从小到大的顺序
	prime := make([]int, 0)

	// 从 2 开始遍历每个数
	for i := 2; i <= n; i++ {
		// 如果 i 没有被标记为合数，说明 i 是质数
		if !notPrime[i] {
			prime = append(prime, i)
		}

		// 遍历当前已知的质数列表，用来筛掉合数
		for _, p := range prime {
			// 如果 i*p 超过了范围 n，后续的质数更大，i*p' 也必然超范围，直接跳出内层循环
			if i*p > n {
				break
			}

			// 将 i*p 标记为合数。因为 p 是质数，且从小到大遍历，
			// 所以 p 是此时 i*p 的最小质因子。
			notPrime[i*p] = true

			// 关键步骤：
			// 如果 i 能被 p 整除，说明 p 是 i 的最小质因子。
			if i%p == 0 {
				// 任何合数 i*p' (其中 p' 是比 p 更大的质数) 的最小质因子也是 p。
				// 为了保证每个合数只被其最小质因子筛掉一次，这里必须 break。
				break
			}
		}
	}
	return prime
}

// CalPrime 计算一个数 n 的所有质因子的个数。
// 例如 n = 12 (2^2 * 3^1)，返回 map[2:2, 3:1]
// 算法思想：
// 1. 遍历从 2 到 sqrt(n) 的所有数 p。
// 2. 如果 p 是 n 的因子，则进入一个循环，持续用 p 去除 n，并计数，直到 n 不能再被 p 整除。
// 3. 遍历结束后，如果 n 大于 1，说明剩下的 n 本身也是一个质因子。
func CalPrime(n int) map[int]int {
	// ks 用来存储每个质因子及其对应的个数
	ks := make(map[int]int)
	// 遍历潜在的质因子 p
	for p := 2; p*p <= n; p++ {
		// 如果 p 是 n 的因子
		if n%p == 0 {
			// 持续用 p 去除 n，并统计 p 的个数
			for n%p == 0 {
				ks[p]++
				n = n / p
			}
		}
	}
	// 如果 n 经过上述操作后仍然大于 1，
	// 说明剩下的 n 是一个大于 sqrt(原始n) 的质因子
	if n > 1 {
		ks[n] = 1
	}
	return ks
}

// Comb 使用动态规划计算组合数 C(n,k)
// 基于帕斯卡三角形（或杨辉三角）的递推关系：C(n,k) = C(n-1,k-1) + C(n-1,k)
// 结果对 mod 取模。
// @param n: 总元素个数
// @param k: 选择的元素个数
// @param mod: 取模的值，用于防止结果溢出
// @return: 返回一个二维数组 ans，其中 ans[i][j] 表示 C(i,j) 的值。最终需要的结果是 ans[n][k]。
func Comb(n, k int, mod int) [][]int {
	// 创建二维数组存储组合数结果
	// ans[i][j] 表示从 i 个元素中选择 j 个元素的组合数 C(i,j)
	ans := make([][]int, n+1)
	for i := range ans {
		ans[i] = make([]int, k+1)
	}

	// 基础情况：C(0,0) = 1
	ans[0][0] = 1

	// 使用递推公式填充表格
	for i := 1; i <= n; i++ {
		// 基础情况：C(i,0) = 1 (从 i 个元素中选 0 个，只有一种方法)
		ans[i][0] = 1
		// j 从 1 到 k，并且 j 必须小于等于 i (因为不能选比总数还多的元素)
		for j := 1; j <= k && j <= i; j++ {
			// 帕斯卡公式: C(i,j) = C(i-1,j-1) + C(i-1,j)
			// 含义：从 i 个元素中选 j 个的方法数 =
			//      (选择第 i 个元素，再从前 i-1 个中选 j-1 个) + (不选择第 i 个元素，从前 i-1 个中选 j 个)
			ans[i][j] = (ans[i-1][j-1] + ans[i-1][j]) % mod
		}
	}
	return ans
}
