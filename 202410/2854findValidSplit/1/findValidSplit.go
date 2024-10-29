package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(findValidSplit([]int{4, 7, 8, 15, 3, 5}))
	fmt.Println(findValidSplit([]int{4, 7, 15, 8, 3, 5}))
}

// 超时了
func findValidSplit(nums []int) int {
	n := len(nums)
	cnt := make([][]int, n)
	mx := slices.Max(nums)
	// mx := int(math.Pow10(6))
	prim := findPrim(mx)

	for i, ch := range nums {
		cnt[i] = find(ch, prim)
	}

	left := make(map[int]int)
	right := make(map[int]int)
	for i := range nums {
		for _, pri := range cnt[i] {
			right[pri]++
		}
	}
	for i := 0; i < n-1; i++ {
		for _, pri := range cnt[i] {
			left[pri]++
			right[pri]--
		}
		if !hasMix(left, right) {
			return i
		}
	}

	return -1
}

func hasMix(left, right map[int]int) bool {
	for k, v := range left {
		if v == 0 {
			continue
		}
		if right[k] > 0 {
			return true
		}
	}
	for k, v := range right {
		if v == 0 {
			continue
		}
		if left[k] > 0 {
			return true
		}
	}
	return false
}

func find(a int, prim []int) []int {
	ans := make([]int, 0)
	for _, ch := range prim {
		if a%ch == 0 {
			ans = append(ans, ch)
		}
	}
	return ans
}

func findPrim(a int) []int {
	prim := make([]bool, a+1)
	for i := 2; i*i <= a; i++ {
		for j := 2; i*j <= a; j++ {
			prim[i*j] = true
		}
	}
	ans := make([]int, 0)
	for i := 2; i <= a; i++ {
		if !prim[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
