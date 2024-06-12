package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSpeedOnTime([]int{1, 3, 2}, 6))
	fmt.Println(minSpeedOnTime([]int{1, 3, 2}, 2.7))
	fmt.Println(minSpeedOnTime([]int{1, 3, 2}, 1.9))
	fmt.Println(minSpeedOnTime([]int{1, 1}, 1.0))
	fmt.Println(minSpeedOnTime([]int{1, 1, 100000}, 2.01))
}

func minSpeedOnTime(dist []int, hour float64) int {
	// mx := slices.Max(dist)
	sum := 0
	for _, ch := range dist {
		sum += ch
	}
	le, ri := 1, int(1e7+1)
	for le < ri {
		mid := le + (ri-le)/2
		if le >= 0 && le <= math.MaxInt && f(dist, mid) > hour {
			le = mid + 1
		} else {
			ri = mid
		}
	}

	if le < 1 || f(dist, le) > hour {
		return -1
	}
	return le
}

// 表示速度是 p 时，到达数组最后需要的时间
func f(nums []int, p int) float64 {
	n := len(nums)
	var ans float64
	for i := 0; i < n-1; i++ {
		ans += float64((nums[i] + p - 1) / p)
	}
	ans += float64(nums[n-1]) / float64(p)
	return ans
}
