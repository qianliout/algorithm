package main

import (
	"fmt"
	"math"
)

// 判断是否为质数（优化版本）
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 解法1：模拟过程（适用于较小的n）
func lastRemainingSimulate(n int) int {
	if n <= 1 {
		return n
	}

	// 创建初始数组 [1, 2, 3, ..., n]
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	// 不断淘汰非质数下标的元素
	for len(arr) > 1 {
		newArr := []int{}
		// 只保留下标为质数的元素（注意：下标从1开始）
		for i := 0; i < len(arr); i++ {
			if isPrime(i + 1) {
				newArr = append(newArr, arr[i])
			}
		}
		arr = newArr
	}

	return arr[0]
}

// 解法2：数学分析（高效方法）
func lastRemainingMath(n int) int {
	if n <= 1 {
		return n
	}

	// 找出所有小于等于n的质数
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	// 问题的解就是质数列表中的最后一个质数
	// 因为每次淘汰后，剩下的都是质数下标的元素
	// 经过多轮筛选，最终剩下的就是最大的满足条件的质数
	return primes[len(primes)-1]
}

// 解法3：递归思路（更直观的理解）
func lastRemainingRecursive(n int) int {
	if n == 1 {
		return 1
	}

	// 找出当前轮次中所有质数下标对应的数字
	// 这些数字会进入下一轮
	nextRound := []int{}
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			nextRound = append(nextRound, i)
		}
	}

	// 如果只剩一个元素，返回它
	if len(nextRound) == 1 {
		return nextRound[0]
	}

	// 否则，递归处理下一轮
	// 注意：下一轮中数组大小变为len(nextRound)
	return lastRemainingRecursive(len(nextRound))
}

func main() {
	testCases := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30}

	fmt.Println("测试结果:")
	fmt.Println("n\t模拟法\t数学法\t递归法")
	fmt.Println("--------------------------------")

	for _, n := range testCases {
		result1 := lastRemainingSimulate(n)
		result2 := lastRemainingMath(n)
		result3 := lastRemainingRecursive(n)

		fmt.Printf("%d\t%d\t%d\t%d\n", n, result1, result2, result3)
	}

	// 测试大一点的数字
	fmt.Println("\n较大数字测试:")
	largeTest := []int{50, 100, 200}
	for _, n := range largeTest {
		// 对于大数字，使用数学方法（模拟法会很慢）
		result := lastRemainingMath(n)
		fmt.Printf("n=%d, 最后剩下的数字: %d\n", n, result)
	}
}
