package main

import (
	"fmt"
)

func main() {
	fmt.Println(splitMessage("this is really a very awesome message", 9))
}

// 分超时
func splitMessage1(message string, limit int) []string {
	// n := len(message)
	// 枚举分隔长度
	for i := 1; ; i++ {
		// 计算当前分段数量下的最大容量
		totalCapacity := 0
		for j := 1; j <= i; j++ {
			tail := fmt.Sprintf("<%d/%d>", j, i)
			messageLength := limit - len(tail)
			if messageLength <= 0 {
				return nil // 如果消息长度不足以容纳尾部信息，则无法分割
			}
			totalCapacity += messageLength
		}

		// 如果当前总容量小于消息长度，继续增加分段数量
		if totalCapacity < len(message) {
			continue
		}

		// 构建分割后的消息
		ans := make([]string, i)
		start := 0
		for j := 1; j <= i; j++ {
			tail := fmt.Sprintf("<%d/%d>", j, i)
			messageLength := limit - len(tail)
			end := start + messageLength
			if end > len(message) {
				ans[j-1] = message[start:] + tail
			} else {
				ans[j-1] = message[start:end] + tail
			}
			start = end
		}

		return ans
	}
}

func splitMessage(message string, limit int) []string {
	n := len(message)
	cp := 0
	tailLength := 0
	i := 1 // 分隔的总份数
	for ; i <= n; i++ {
		if i >= 1 && i < 10 {
			tailLength = 5 // <1/1>
		}
		if i >= 10 && i < 100 {
			if i == 10 {
				cp -= 9 // 前面9个
			}
			tailLength = 7 // <10/10>
		}
		if i >= 100 && i < 1000 {
			if i == 100 {
				cp -= 99
			}
			tailLength = 9
		}
		if i >= 1000 && i < 10000 {
			if i == 1000 {
				cp -= 999
			}
			tailLength = 11
		}
		if tailLength >= limit {
			return nil
		}
		cp += limit - tailLength
		if cp >= n {
			break
		}
	}
	// 求出了总数了，组装结果
	ans := make([]string, i)
	for j := range ans {
		t := fmt.Sprintf("<%d/%d>", j+1, i)
		if j == i-1 {
			ans[j] = message + t
			break
		}
		m := limit - len(t)
		ans[j] = message[:m] + t
		message = message[m:]
	}

	return ans
}
