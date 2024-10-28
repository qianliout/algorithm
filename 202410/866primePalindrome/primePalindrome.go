package main

func main() {

}

func primePalindrome(n int) int {
	// 不超过 11 的质数都是回文数，因此当 n≤11 时，大于等于 n 的最小回文质数即为大于等于 n 的最小质数且不超过 11，可以从 n 开始从小到大枚举每个整数，寻找最小质数。
	if n <= 11 {
		num := n
		for !prime(num) {
			num++
		}
		return num
	}

	ans := 0
	// 用 left 表示回文数的左侧一半数位包含中间数位，即回文数的位数等于 left 的位数的两倍减 1。例如，left=10 对应的回文数是 101，left=1245 对应的回文数是 1245421
	left := 10
	for {
		p := generatePalindrome(left)
		if p >= n && prime(p) {
			ans = p
			break
		} else {
			left++
		}
	}
	return ans
}

// 判断一个数是不是质数
func prime(n int) bool {
	if n <= 1 {
		return false
	}
	if n&1 == 0 {
		return n == 2
	}
	for i := 3; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func generatePalindrome(left int) int {
	pr := left
	left = left / 10
	for left > 0 {
		pr = pr*10 + left%10
		left = left / 10
	}
	return pr
}
