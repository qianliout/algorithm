package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(partitionArray([]int{16, 8, 17, 0, 3, 17, 8, 20}, 10))
}

// 直接贪心算法会出错：
// 需要直接排序
// fmt.Println(partitionArray([]int{16, 8, 17, 0, 3, 17, 8, 20}, 10))
func partitionArray1(nums []int, k int) int {
	sort.Ints(nums)
	st := make([]pair, 0)
	for _, ch := range nums {
		flag := false
		for j := len(st) - 1; j >= 0; j-- {
			p, b := add(st[j], ch, k)
			if b {
				flag = true
				st[j] = p
				break
			}
		}

		if !flag {
			st = append(st, pair{Mx: ch, Mi: ch})
		}
	}
	return len(st)
}

type pair struct {
	Mx, Mi int
}

func add(pa pair, num, k int) (pair, bool) {
	nw := pair{
		Mx: max(pa.Mx, num),
		Mi: min(pa.Mi, num),
	}
	if nw.Mx-nw.Mi <= k {
		return nw, true
	}
	return pa, false
}

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	mi := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-mi > k {
			ans++
			mi = nums[i]
		}
	}
	return ans
}
