package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// fmt.Println(largestOverlap([][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}))
	fmt.Println(largestOverlap([][]int{{0, 0, 0, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		[][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}}))
}

// 没有搞定
func largestOverlap(img1 [][]int, img2 [][]int) int {
	m := len(img1)
	img11 := make([][]int, m)
	img22 := make([][]int, m)
	for i := range img11 {
		img11[i] = append([]int{}, img1[i]...)
		img22[i] = append([]int{}, img2[i]...)
	}

	return max(help(img1, img2), help(img22, img11))
}

func help(img1 [][]int, img2 [][]int) int {
	m := len(img1)
	image1 := make([]int, 0)
	image2 := make([]int, 0)
	for i := range img2 {
		image1 = append(image1, gen(img1[i]))
		image2 = append(image2, gen(img2[i]))
	}
	ans := 0
	for i := 0; i < m; i++ {
		// 左移
		moved := moveLeft(image1, i)
		// 向下移动
		for j := 0; j < m; j++ {
			ans = max(ans, cal(moved, image2[j:]))
		}
		// 向上移动
		for j := 0; j < m; j++ {
			ans = max(ans, cal(moved, image2[:m-j]))
		}
	}
	return ans
}

// 一行的二进制集合
func gen(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			ans = ans | (1 << (n - i - 1))
		}
	}
	return ans
}

func moveLeft(source []int, step int) []int {
	ans := make([]int, len(source))

	for i, ch := range source {
		ans[i] = ch >> step
	}
	return ans
}

func moveRight(source []int, step int) []int {
	ans := make([]int, len(source))

	for i, ch := range source {
		ans[i] = ch << step
	}
	return ans
}

func cal(source []int, target []int) int {
	ans := 0
	n := len(target)
	for i := 0; i < n; i++ {
		ans += bits.OnesCount(uint(source[i] & target[i]))
	}
	return ans
}
