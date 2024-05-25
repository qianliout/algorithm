package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// fmt.Println(maximumScore([]int{19, 12, 14, 6, 10, 18}, 3))
	// fmt.Println(maximumScore([]int{8, 3, 9, 3, 8}, 2))
	fmt.Println(maximumScore([]int{3289, 2832, 14858, 22011}, 6)) // 256720975

}

func maximumScore(nums []int, k int) int {
	base := int(math.Pow(10, 9)) + 7
	mem := make(map[int]int)
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = -1
		right[i] = n
	}
	primes := genPrimes()
	ppp := make([]int, n)
	for i := range ppp {
		ppp[i] = primes[nums[i]]
	}
	// 向右
	st := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(st) > 0 && primes[nums[i]] > primes[nums[st[len(st)-1]]] {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && primes[nums[i]] >= primes[nums[st[len(st)-1]]] {
			left[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool { return nums[ids[i]] > nums[ids[j]] })

	ans := 1
	for _, i := range ids {
		// 经过上面的排序，此时 i就是先选取nums中最大元素的下标
		tot := (i - left[i]) * (right[i] - i)
		if tot >= k {
			ans = (ans * pow(nums[i], k, base, mem)) % base
			break
		}
		ans = (ans * pow(nums[i], tot, base, mem)) % base
		// 这里为啥要减 tot 呢，
		// 因为 这里 tot 是指以nums[i]做为子数组的最大元素所选取的子数组的个数
		k -= tot // 更新剩余操作次数
	}
	return ans
}

func genPrimes() []int {
	mx := int(math.Pow(10, 5))
	// primes[i]=k 表示 i 这个数可以分解成 k个不同的质数
	primes := make([]int, mx+1)
	for i := 2; i <= mx; i++ {
		if primes[i] == 0 {
			for j := i; j <= mx; j += i {
				primes[j]++ // i 是 j 的一个质因子
			}
		}
	}
	return primes
}

func pow(x, n int, base int, mem map[int]int) int {
	if va, ok := mem[n]; ok {
		return va
	}
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % base
		}
		x = x * x % base
	}
	return res
}
