package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(bestRotation([]int{2, 3, 1, 4, 0}))
}

func categorizeBox(length int, width int, height int, mass int) string {
	p4 := int(math.Pow10(4))

	b := length >= p4 || width >= p4 || height >= p4 || (length*width*height) >= int(math.Pow10(9))
	h := mass >= 100
	if b && h {
		return "Both"
	}
	if !b && !h {
		return "Neither"
	}
	if b && !h {
		return "Bulky"
	}
	if !b && h {
		return "Heavy"
	}
	return ""
}

func bestRotation(nums []int) int {
	n := len(nums)
	d := make([]int, int(math.Pow10(5)+1))
	for i := range nums {
		// 这里计算上下限是难点，
		// 参考三叶的题解
		a := (i - (n - 1) + n) % n
		b := (i - nums[i] + n) % n
		if a <= b {
			d[a]++
			d[b+1]--
		} else {
			d[0]++
			d[b+1]--

			d[a]++
			d[n]--
		}
	}

	for i := 1; i <= n; i++ {
		d[i] += d[i-1]
	}
	ans := 0
	for i := 0; i <= n; i++ {
		if d[i] > d[ans] {
			ans = i
		}
	}
	return ans
}

// https://leetcode.cn/problems/smallest-rotation-with-highest-score/solutions/1322229/gong-shui-san-xie-shang-xia-jie-fen-xi-c-p6kh/
