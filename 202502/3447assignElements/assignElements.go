package main

import (
	"slices"
)

func main() {

}

func assignElements(groups []int, elements []int) []int {
	mx := slices.Max(groups)
	target := make([]int, mx+1)
	for i := range target {
		target[i] = -1 // 默认值是-1
	}
	for i, ch := range elements {
		if ch > mx || target[ch] >= 0 {
			// 剪枝，当枚举2之后，就不用再枚举4，6，8等
			continue
		}
		for j := ch; j <= mx; j = j + ch {
			if target[j] < 0 { // 第一次出现
				target[j] = i
			}
		}
	}
	// 输出答案
	ans := make([]int, len(groups))
	for i, ch := range groups {
		ans[i] = target[ch]
	}
	return ans
}
