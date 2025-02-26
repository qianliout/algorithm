package main

import (
	"fmt"
)

func main() {
	fmt.Println(fileCombination(18))
}

// 滑动窗口的做法
func fileCombination3(target int) [][]int {
	win := 0
	ans := make([][]int, 0)
	for le, ri := 1, 1; ri <= target; ri++ {
		win += ri
		// 维护
		for win > target {
			win -= le
			le++
		}
		if win == target {
			lev := make([]int, 0)
			for j := le; j <= ri; j++ {
				lev = append(lev, j)
			}
			if len(lev) >= 2 {
				ans = append(ans, lev)
			}
		}
	}
	return ans
}

// 二分的方式
func fileCombination2(target int) [][]int {
	ans := make([][]int, 0)
	sum := make([]int, target+1)
	for i := 1; i <= target; i++ {
		sum[i] = sum[i-1] + i
	}
	find := func(ch int) (int, bool) {
		le, ri := 0, target
		for le < ri {
			// 找>=ch 的左端点
			mid := le + (ri-le)/2
			if mid < target && sum[mid] >= ch {
				ri = mid
			} else {
				le = mid + 1
			}
		}
		if le >= 0 && le < target && sum[le] == ch {
			return le, true
		}
		return -1, false
	}

	for i := 1; i <= target; i++ {
		nw := sum[i]
		start, b := find(nw - target)
		if b {
			lev := make([]int, 0)
			for j := start + 1; j <= i; j++ {
				lev = append(lev, j)
			}
			if len(lev) >= 2 {
				ans = append(ans, lev)
			}
		}
	}
	return ans
}

// 滑动窗口的做法,这样做不能得到正确的结果
func fileCombination(target int) [][]int {
	win := 0
	ans := make([][]int, 0)
	for le, ri := 1, 1; ri <= target; ri++ {
		win += ri
		if win == target {
			lev := make([]int, 0)
			for j := le; j <= ri; j++ {
				lev = append(lev, j)
			}
			if len(lev) >= 2 {
				ans = append(ans, lev)
			}
		}
		for win > target {
			win -= le
			le++
		}
	}
	return ans
}
