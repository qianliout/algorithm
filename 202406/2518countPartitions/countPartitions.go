package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(countPartitions([]int{1, 2, 3, 4}, 4))
	fmt.Println(countPartitions([]int{790, 555, 729, 447, 538, 657, 258, 716, 645, 349, 148, 860, 425, 401, 282, 889, 309, 720, 228, 39, 366, 107, 765, 546, 791, 938, 154, 85, 845, 656}, 558))
	fmt.Println(countPartitions([]int{478, 721, 51, 352, 361, 66, 22, 807, 59, 275, 114, 169, 855, 103, 509, 592, 769, 384, 670, 764, 382, 466, 69, 787, 69, 217, 992, 37, 805, 842, 760, 515, 442, 77, 660, 449, 471, 752, 743, 947, 616, 246, 46, 973, 860, 264, 852, 675, 139, 60, 368, 554, 723, 350, 870, 710, 966, 633, 99, 146, 175, 659, 941, 592, 614, 221, 775, 76, 228, 372, 430, 521, 248, 591, 683, 600, 439, 58, 653, 810, 206, 488, 714, 531, 383, 251, 566, 448, 580, 627, 689, 818, 525, 321, 127, 499, 440}, 634))
}

// 坏分区，小于 k
func countPartitions(nums []int, k int) int {
	mod := int(math.Pow10(9)) + 7
	sum := 0
	for _, ch := range nums {
		sum += ch
	}
	if sum < 2*k {
		return 0
	}
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	// 初值 // 空集合
	for j := 0; j < k; j++ {
		dp[0][j] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < k; j++ {
			// 不选
			no := dp[i-1][j]
			// 选
			yes := 0
			if j-nums[i-1] >= 0 {
				yes = dp[i-1][j-nums[i-1]]
			}
			dp[i][j] = (no + yes) % mod
		}
	}
	// 最后这一步是最容易出错的，以后只要是用到 mod 的题，都要这样做
	return ((pow(2, n, mod)-dp[n][k-1]*2)%mod + mod) % mod
}

func pow(x, n int, mod int) int {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		a := pow(x, n/2, mod)
		return a * a % mod
	}
	return (x * pow(x, n-1, mod)) % mod
}

func cal(n int) int {
	if n == 1 {
		return 1
	}
	return n * cal(n-1)
}

func getCommon(nums1 []int, nums2 []int) int {
	cnt := make(map[int]int)
	for _, ch := range nums2 {
		cnt[ch]++
	}
	sort.Ints(nums1)
	for _, ch := range nums1 {
		if cnt[ch] > 0 {
			return ch
		}
	}
	return -1
}
