package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumScore([]int{3289, 2832, 14858, 22011}, 6))
}

const mod int = 1e9 + 7

// 预处理 omega

const mx int = 1e5

var omega [mx + 1]int

func init() {

	for i := 2; i <= mx; i++ {

		if omega[i] == 0 { // i 是质数

			for j := i; j <= mx; j += i {

				omega[j]++ // i 是 j 的一个质因子

			}

		}

	}

}

func maximumScore(nums []int, k int) int {

	n := len(nums)

	left := make([]int, n) // 质数分数 >= omega[nums[i]] 的左侧最近元素下标

	right := make([]int, n) // 质数分数 >  omega[nums[i]] 的右侧最近元素下标

	for i := range right {

		right[i] = n

	}

	ppp := make([]int, n)
	for i := range ppp {
		ppp[i] = omega[nums[i]]
	}
	fmt.Println(ppp)

	st := []int{-1}

	for i, v := range nums {

		for len(st) > 1 && omega[nums[st[len(st)-1]]] < omega[v] {

			right[st[len(st)-1]] = i

			st = st[:len(st)-1]

		}

		left[i] = st[len(st)-1]

		st = append(st, i)

	}

	id := make([]int, n)
	fmt.Println("left", left, "right", right)

	for i := range id {

		id[i] = i

	}

	sort.Slice(id, func(i, j int) bool { return nums[id[i]] > nums[id[j]] })
	fmt.Println("ids", id)
	ans := 1

	for _, i := range id {

		tot := (i - left[i]) * (right[i] - i)

		if tot >= k {

			ans = ans * pow(nums[i], k) % mod

			break

		}

		ans = ans * pow(nums[i], tot) % mod

		k -= tot // 更新剩余操作次数

	}

	return ans

}

func pow(x, n int) int {

	res := 1

	for ; n > 0; n /= 2 {

		if n%2 > 0 {

			res = res * x % mod

		}

		x = x * x % mod

	}

	return res

}
