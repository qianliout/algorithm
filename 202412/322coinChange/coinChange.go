package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

// 会超时
func coinChange1(coins []int, amount int) int {
	cnt := make(map[int]bool)
	for _, ch := range coins {
		cnt[ch] = true
	}
	inf := 1 << 31
	f := make([]int, amount+1)
	for i := range f {
		f[i] = inf
	}
	f[0] = 0
	for i := 0; i <= amount; i++ {
		for j := i; j >= 0; j-- {
			if cnt[i-j] && f[i] > f[j]+1 {
				f[i] = f[j] + 1
			}
		}
	}
	if f[amount] == inf {
		return -1
	}
	return f[amount]
}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	inf := amount + 1
	f := make([]int, amount+1)
	for i := range f {
		f[i] = inf
	}
	f[0] = 0
	for i := 0; i <= amount; i++ {
		for _, ch := range coins {
			if i < ch {
				break
			}
			f[i] = min(f[i], f[i-ch]+1)
		}
	}
	if f[amount] == inf {
		return -1
	}
	return f[amount]
}
