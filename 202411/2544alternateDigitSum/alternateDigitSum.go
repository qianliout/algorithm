package main

func main() {

}

func alternateDigitSum(n int) int {
	sign := 1
	ans := 0
	for n > 0 {
		ans += sign * (n % 10)
		n /= 10
		sign *= -1
	}

	// 如果n中有偶数个数位，那么第一位是正，最后一位是负，按上面的算法刚好算反了，奇数位是一个道理
	// 可以用 11 1 两个数去思考
	ans = ans * -sign
	return ans
}
