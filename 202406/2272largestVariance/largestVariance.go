package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestVariance("aababbb"))
	fmt.Println(largestVariance("aabbbbaa"))
}

// 这样写是对的，但是我没有太理解
func largestVariance1(s string) int {
	ans := 0
	// 建设最大的字符是 a，最小的字符是 b 分别统计
	// 因为只有 a 和 b 都存在时，才能算答案，所以要用 diff 和diff2B两个变量
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if a == b {
				continue
			}

			diff, diff2B := 0, -len(s)
			// diff2B = -len(s)表示还没有找到 b
			for _, ch := range s {
				if ch == a {
					diff++
					diff2B++
				} else if ch == b {
					diff--
					diff2B = diff
					diff = max(diff, 0) // 这一步很重要，但是为啥 diff2b不加判断呢
				}
				ans = max(ans, diff2B)
			}
		}
	}
	return ans
}

func largestVariance(s string) int {
	ans := 0
	// 建设最大的字符是 a，最小的字符是 b 分别统计
	// 因为只有 a 和 b 都存在时，才能算答案，所以要用 diff 和diff2B两个变量
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if a == b {
				continue
			}
			cnta, cntb := 0, 0
			// diff2B = -len(s)表示还没有找到 b
			for _, ch := range s {
				if ch == a {
					cnta++
				} else if ch == b {
					cntb++
				}

				if cnta > 0 && cntb > 0 {
					ans = max(ans, cnta-cntb)
				}
			}
		}
	}
	return ans
}
